import { Box, Text, Image, Badge, HStack, Icon, Flex } from "@chakra-ui/react";
import { Card } from "@/lib/types/cards.types";
import { FiCheckCircle, FiCheckSquare, FiCircle, FiEdit2, FiAlignLeft } from "react-icons/fi";
import { useState, useRef } from "react";
import { updateCard } from "@/lib/services/cards";
import { toaster } from "@/components/ui/toaster";
import { Tooltip } from "@/components/ui/tooltip";
import { useBoardStore } from "@/lib/store/board";
import CardQuickEdit from "./CardQuickEdit";
import { useSortable } from "@dnd-kit/sortable";
import { CSS } from "@dnd-kit/utilities";

interface BoardCardProps {
    card: Card;
    listId: string;
    boardId: string;
    onUpdate: () => void;
    onClick: () => void;
}

export default function BoardCard({ card, listId, boardId, onUpdate, onClick }: BoardCardProps) {
    const { metadata, cards } = useBoardStore();
    const storeCard = cards.find(c => c.id === card.id) || card;

    const {
        attributes,
        listeners,
        setNodeRef,
        transform,
        transition,
        isDragging
    } = useSortable({ id: storeCard.id });
    const style = {
        transition,
        transform: CSS.Translate.toString(transform),
        opacity: isDragging ? 0.5 : 1,
    };

    const [isHovered, setIsHovered] = useState(false);
    const [isEditing, setIsEditing] = useState(false);
    const cardRef = useRef<HTMLDivElement>(null);
    const [editPos, setEditPos] = useState<{ top: number; left: number; width: number } | null>(null);

    const handleToggleComplete = async (e: React.MouseEvent) => {
        e.stopPropagation();
        try {
            await updateCard({
                cardID: storeCard.id,
                listID: listId,
                boardID: boardId,
                completed: !storeCard.completed
            });
            onUpdate();
        } catch (error) {
            toaster.create({
                title: "Failed to update card",
                type: "error",
            });
        }
    };

    const handleEditClick = (e: React.MouseEvent) => {
        e.stopPropagation();
        if (cardRef.current) {
            const rect = cardRef.current.getBoundingClientRect();
            setEditPos({
                top: rect.top,
                left: rect.left,
                width: rect.width
            });
            setIsEditing(true);
            setIsHovered(false); // Hide hover effects when editing
        }
    };

    const handleSaveTitle = async (newTitle: string) => {
        try {
            await updateCard({
                cardID: storeCard.id,
                listID: listId,
                boardID: boardId,
                title: newTitle
            });
            onUpdate();
            setIsEditing(false);
        } catch (error) {
            toaster.create({
                title: "Failed to update card title",
                type: "error",
            });
        }
    };

    const displayMembers = (storeCard.member_ids || []).map(id => {
        const member = metadata?.members.find(m => m.id === id);
        return member ? {
            userID: member.id,
            username: member.username,
            fullName: member.fullName,
            email: member.email,
            avatar: "",
            isCardMember: true
        } : null;
    }).filter(Boolean) as any[];

    return (
        <div ref={setNodeRef} {...attributes} style={style} {...listeners}>
            <Box
                ref={cardRef}
                borderRadius="md"
                boxShadow="sm"
                mb={2}
                cursor="pointer"
                height="auto"
                minHeight="0"
                _hover={{
                    bg: storeCard.coverSize === "full" && storeCard.cover
                        ? (storeCard.cover.startsWith("#") ? storeCard.cover : "gray.700")
                        : "gray.600",
                    boxShadow: "md",
                    borderColor: "blue.400"
                }}
                border="1px solid"
                borderColor="whiteAlpha.100"
                transition="all 0.5s ease-in-out"
                onMouseEnter={() => setIsHovered(true)}
                onMouseLeave={() => setIsHovered(false)}
                role="group"
                opacity={isEditing ? 0 : 1}
                onClick={onClick}
                position="relative"
                overflow="hidden"
                bg={storeCard.coverSize === "full" && storeCard.cover && storeCard.cover.startsWith("#") ? storeCard.cover : "gray.700"}
            >
                {storeCard.coverSize === "full" && storeCard.cover && !storeCard.cover.startsWith("#") && (
                    <Image
                        src={storeCard.cover}
                        w="full"
                        h="auto"
                        display="block"
                        alt="Cover"
                    />
                )}
                {storeCard.cover && storeCard.coverSize !== "full" && (
                    storeCard.cover.startsWith("#") ? (
                        <Box
                            h={10}
                            w="full"
                            bg={storeCard.cover}
                            borderTopRadius="md"
                            mb={2}
                        />
                    ) : (
                        <Image
                            src={storeCard.cover}
                            w="full"
                            h="auto"
                            display="block"
                            alt="Cover"
                            borderTopRadius="md"
                            mb={2}
                        />
                    )
                )}
                {storeCard.coverSize === "full" && storeCard.cover && !storeCard.cover.startsWith("#") && (
                    <Box
                        position="absolute"
                        bottom={0}
                        left={0}
                        w="full"
                        h="full"
                        bgGradient="linear(to-t, blackAlpha.900, transparent)"
                        pointerEvents="none"
                    />
                )}

                {storeCard.coverSize != "full" && storeCard.labels && storeCard.labels.length > 0 && (
                    <Flex
                        gap={1}
                        px={2}
                        pt={storeCard.cover && storeCard.coverSize !== "full" ? 0 : 2}
                        wrap="wrap"
                        mb={1}
                        position={storeCard.coverSize === "full" && storeCard.cover && !storeCard.cover.startsWith("#") ? "absolute" : "relative"}
                        top={storeCard.coverSize === "full" && storeCard.cover && !storeCard.cover.startsWith("#") ? 0 : undefined}
                        left={storeCard.coverSize === "full" && storeCard.cover && !storeCard.cover.startsWith("#") ? 0 : undefined}
                        zIndex={1}
                    >
                        {storeCard.labels.map((label: any) => (
                            <Box
                                key={label.labelID || label.id}
                                bg={label.color}
                                w={10}
                                h={2}
                                borderRadius="full"
                                title={label.name}
                            />
                        ))}
                    </Flex>
                )}

                <Flex
                    justify={storeCard.coverSize === "full" && storeCard.cover ? "flex-end" : "start"}
                    align="start"
                    gap={2}
                    p={2}
                    h={storeCard.coverSize === "full" && storeCard.cover && !storeCard.cover.startsWith("#") ? "100%" : "auto"}
                    position={storeCard.coverSize === "full" && storeCard.cover && !storeCard.cover.startsWith("#") ? "absolute" : "relative"}
                    bottom={0}
                    left={0}
                    right={0}
                    zIndex={1}
                    direction="column"
                >
                    <Flex w="full" gap={2} align="start">
                        {isHovered && !storeCard.completed && (
                            <Icon
                                as={FiCircle}
                                color={storeCard.coverSize === "full" && storeCard.cover ? "whiteAlpha.800" : "gray.400"}
                                boxSize={4}
                                mt={0.5}
                                animation="fadeIn 0.3s ease-out 0.1s both"
                                onClick={handleToggleComplete}
                                cursor="pointer"
                                _hover={{ color: "green.400" }}
                            />
                        )}
                        {storeCard.completed && (
                            <Icon
                                as={FiCheckCircle}
                                color="green.400"
                                boxSize={4}
                                mt={0.5}
                                animation="fadeIn 0.3s ease-out 0.1s both"
                                onClick={handleToggleComplete}
                                cursor="pointer"
                            />
                        )}
                        <Text
                            fontSize={storeCard.coverSize === "full" && storeCard.cover && !storeCard.cover.startsWith("#") ? "md" : "sm"}
                            fontWeight={storeCard.coverSize === "full" && storeCard.cover && !storeCard.cover.startsWith("#") ? "bold" : "normal"}
                            color="white"
                            flex={1}
                            textDecoration={storeCard.completed ? "line-through" : "none"}
                            opacity={storeCard.completed ? 0.7 : 1}
                            animation={isHovered ? "slideIn 0.4s ease-out 0.2s both" : "none"}
                            textShadow={storeCard.coverSize === "full" && storeCard.cover && !storeCard.cover.startsWith("#") ? "0 1px 2px rgba(0,0,0,0.8)" : "none"}
                        >
                            {storeCard.title}
                        </Text>
                        {isHovered && (
                            <Icon
                                as={FiEdit2}
                                color={storeCard.coverSize === "full" && storeCard.cover ? "whiteAlpha.800" : "gray.400"}
                                boxSize={3.5}
                                _hover={{ color: "white" }}
                                mt={0.5}
                                animation="fadeIn 0.3s ease-out 0.3s both"
                                onClick={handleEditClick}
                            />
                        )}
                    </Flex>

                    {(storeCard.description || displayMembers.length > 0) && storeCard.coverSize !== "full" && (
                        <Flex w="full" justify="space-between" align="center" mt={2}>
                            <Flex gap={3} align="center">
                                {storeCard.description && (
                                    <Tooltip content="This card has a description" positioning={{ placement: "bottom" }}>
                                        <Icon as={FiAlignLeft} color="gray.400" boxSize={4} />
                                    </Tooltip>
                                )}
                            </Flex>

                            {displayMembers.length > 0 && (
                                <Flex gap={1}>
                                    {displayMembers.slice(0, 2).map((member) => (
                                        <Tooltip
                                            key={member.userID}
                                            content={member.fullName || member.username}
                                            positioning={{ placement: "bottom" }}
                                        >
                                            <Box
                                                bg="blue.500"
                                                borderRadius="full"
                                                w={6}
                                                h={6}
                                                display="flex"
                                                alignItems="center"
                                                justifyContent="center"
                                                fontSize="xs"
                                                fontWeight="bold"
                                                color="white"
                                                title={member.username}
                                            >
                                                {member.username.charAt(0).toUpperCase()}
                                            </Box>
                                        </Tooltip>
                                    ))}
                                    {displayMembers.length > 2 && (
                                        <Box
                                            bg="gray.600"
                                            borderRadius="full"
                                            w={6}
                                            h={6}
                                            display="flex"
                                            alignItems="center"
                                            justifyContent="center"
                                            fontSize="xs"
                                            fontWeight="bold"
                                            color="white"
                                        >
                                            +{displayMembers.length - 2}
                                        </Box>
                                    )}
                                </Flex>
                            )}
                        </Flex>
                    )}
                </Flex>

            </Box>

            <CardQuickEdit
                isOpen={isEditing}
                onClose={() => setIsEditing(false)}
                card={storeCard}
                position={editPos}
                onSave={handleSaveTitle}
                onUpdate={onUpdate}
                listId={listId}
                boardId={boardId}
            />
        </div>
    );
}
