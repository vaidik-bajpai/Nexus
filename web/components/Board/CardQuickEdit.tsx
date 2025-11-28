import { Box, Button, Textarea, VStack, HStack, Text, Icon, Portal, Flex } from "@chakra-ui/react";
import { Card } from "@/lib/types/cards.types";
import {
    FiLayout, FiTag, FiUser, FiImage, FiClock,
    FiArrowRight, FiCopy, FiLink, FiArchive, FiMonitor
} from "react-icons/fi";
import { useEffect, useRef, useState } from "react";

interface CardQuickEditProps {
    isOpen: boolean;
    onClose: () => void;
    card: Card;
    position: { top: number; left: number; width: number } | null;
    onSave: (newTitle: string) => void;
}

export default function CardQuickEdit({ isOpen, onClose, card, position, onSave }: CardQuickEditProps) {
    const [title, setTitle] = useState(card.title);
    const contentRef = useRef<HTMLDivElement>(null);

    useEffect(() => {
        setTitle(card.title);
    }, [card]);

    // Close on click outside
    useEffect(() => {
        const handleClickOutside = (event: MouseEvent) => {
            if (contentRef.current && !contentRef.current.contains(event.target as Node)) {
                onClose();
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

    const menuItems = [
        { icon: FiLayout, label: "Open card" },
        { icon: FiTag, label: "Edit labels" },
        { icon: FiUser, label: "Change members" },
        { icon: FiImage, label: "Change cover" },
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
                                    h="120px"
                                    w="full"
                                    bgImage={`url(${card.cover})`}
                                    bgSize="cover"
                                    backgroundPosition="center"
                                    borderTopRadius="md"
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
                            <Button
                                key={index}
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
                        ))}
                    </VStack>
                </Flex>
            </Box>
        </Portal>
    );
}
