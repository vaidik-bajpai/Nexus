import { Box, Button, Textarea, VStack, HStack, Text, Icon, Portal, Flex, Popover } from "@chakra-ui/react";
import { Card, CardMember } from "@/lib/types/cards.types";
import {
    FiLayout, FiTag, FiUser, FiImage, FiClock,
    FiArrowRight, FiCopy, FiLink, FiArchive, FiMonitor
} from "react-icons/fi";
import { useEffect, useRef, useState } from "react";
import ChangeMembers from "./ChangeMembers";
import ChangeCover from "./ChangeCover";
import { updateCard, getCardDetail } from "@/lib/services/cards";
import { toaster } from "@/components/ui/toaster";

interface CardQuickEditProps {
    isOpen: boolean;
    onClose: () => void;
    card: Card;
    position: { top: number; left: number; width: number } | null;
    onSave: (newTitle: string) => void;
    onUpdate: () => void;
    listId: string;
    boardId: string;
}

export default function CardQuickEdit({ isOpen, onClose, card, position, onSave, onUpdate, listId, boardId }: CardQuickEditProps) {
    const [title, setTitle] = useState(card.title);
    const contentRef = useRef<HTMLDivElement>(null);

    useEffect(() => {
        setTitle(card.title);
    }, [card, isOpen, boardId, listId]);

    // Close on click outside
    useEffect(() => {
        const handleClickOutside = (event: MouseEvent) => {
            if (contentRef.current && !contentRef.current.contains(event.target as Node)) {
                const target = event.target as HTMLElement;
                if (target.closest(".chakra-popover__content")) {
                    return;
                }

                if (contentRef.current && !contentRef.current.contains(event.target as Node)) {
                    onClose();
                }
            }
        };

        if (isOpen) {
            document.addEventListener("mousedown", handleClickOutside);
            // Prevent scrolling when open
            document.body.style.overflow = "hidden";
        }

        return () => {
            document.removeEventListener("mousedown", handleClickOutside);
            document.body.style.overflow = "unset";
        };
    }, [isOpen, onClose]);

    if (!isOpen || !position) return null;

    const handleCoverUpdate = async (field: string, value: any) => {
        try {
            await updateCard({
                cardID: card.id,
                listID: listId,
                boardID: boardId,
                [field]: value
            });
            onUpdate();
        } catch (error) {
            toaster.create({
                title: "Failed to update cover",
                type: "error",
            });
        }
    };

    const menuItems = [
        { icon: FiLayout, label: "Open card" },
        { icon: FiTag, label: "Edit labels" },
        { icon: FiUser, label: "Change members", portal: <ChangeMembers memberIds={card.member_ids || []} cardID={card.id} listID={listId} boardID={boardId} /> },
        {
            icon: FiImage,
            label: "Change cover",
            portal: <ChangeCover
                onClose={() => { }} // Popover handles closing
                onUpdate={(cover, coverSize) => {
                    // We need to propagate this up
                    // For now, let's assume we receive a prop or context
                    // See comment above about handleCoverUpdate
                }}
                currentCover={card.cover}
                currentSize={card.coverSize}
            />
        },
        { icon: FiClock, label: "Edit dates" },
        { icon: FiArrowRight, label: "Move" },
        { icon: FiCopy, label: "Copy card" },
        { icon: FiLink, label: "Copy link" },
        { icon: FiMonitor, label: "Mirror" },
        { icon: FiArchive, label: "Archive" },
    ];

    return (
        <Portal>
            <Box
                position="fixed"
                top={0}
                left={0}
                w="100vw"
                h="100vh"
                bg="blackAlpha.600"
                zIndex={1500}
            >
                <Flex
                    ref={contentRef}
                    position="absolute"
                    top={`${position.top}px`}
                    left={`${position.left}px`}
                    align="flex-start"
                    gap={2}
                >
                    {/* Editor Area */}
                    <Box w={`${position.width}px`}>
                        <Box
                            bg="gray.700"
                            borderRadius="md"
                            mb={2}
                        >
                            {card.cover && (
                                <Box
                                    h={card.coverSize === "full" ? "200px" : "120px"} // Adjust height based on size? Or just style.
                                    w="full"
                                    bg={card.cover.startsWith("#") ? card.cover : undefined}
                                    bgImage={!card.cover.startsWith("#") ? `url(${card.cover})` : undefined}
                                    bgSize="cover"
                                    backgroundPosition="center"
                                    borderTopRadius="md"
                                // If full size, maybe we want text over it? 
                                // For quick edit, let's keep it simple for now.
                                />
                            )}
                            <Textarea
                                value={title}
                                onChange={(e) => setTitle(e.target.value)}
                                bg="transparent"
                                border="none"
                                color="white"
                                fontSize="sm"
                                resize="none"
                                p={2}
                                _focus={{ ring: 0 }}
                            />
                        </Box>
                        <Button
                            colorPalette="blue"
                            size="sm"
                            onClick={() => onSave(title)}
                        >
                            Save
                        </Button>
                    </Box>

                    {/* Sidebar Menu */}
                    <VStack align="stretch" gap={1}>
                        {menuItems.map((item, index) => (
                            <Popover.Root key={index} positioning={{ placement: "bottom-start" }}>
                                <Popover.Trigger asChild>
                                    <Button
                                        variant="subtle"
                                        justifyContent="flex-start"
                                        size="xs"
                                        bg="blackAlpha.600"
                                        color="white"
                                        w={"fit-content"}
                                        _hover={{ bg: "blackAlpha.800", transform: "translateX(4px)" }}
                                        transition="all 0.2s"
                                        px={3}
                                        py={2}
                                    >
                                        <Icon as={item.icon} size={"xs"} />
                                        <Text fontSize={"xs"}>{item.label}</Text>
                                    </Button>
                                </Popover.Trigger>
                                {item.label === "Change cover" ? (
                                    <ChangeCover
                                        onClose={() => { }}
                                        onUpdate={handleCoverUpdate}
                                        currentCover={card.cover}
                                        currentSize={card.coverSize}
                                    />
                                ) : item.portal}
                            </Popover.Root>
                        ))}
                    </VStack>
                </Flex>
            </Box>
        </Portal>
    );
}
