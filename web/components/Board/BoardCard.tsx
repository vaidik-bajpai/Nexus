import { Box, Text, Image, Badge, HStack, Icon, Flex } from "@chakra-ui/react";
import { Card } from "@/lib/types/cards.types";
import { FiCheckCircle, FiCheckSquare, FiCircle, FiEdit2 } from "react-icons/fi";
import { useState } from "react";
import { updateCard } from "@/lib/services/cards";
import { toaster } from "@/components/ui/toaster";

interface BoardCardProps {
    card: Card;
    listId: string;
    boardId: string;
    onUpdate: () => void;
}

export default function BoardCard({ card, listId, boardId, onUpdate }: BoardCardProps) {
    const [isHovered, setIsHovered] = useState(false);

    const handleToggleComplete = async (e: React.MouseEvent) => {
        e.stopPropagation();
        try {
            await updateCard({
                cardID: card.id,
                listID: listId,
                boardID: boardId,
                completed: !card.completed
            });
            onUpdate();
        } catch (error) {
            toaster.create({
                title: "Failed to update card",
                type: "error",
            });
        }
    };

    return (
        <Box
            bg="gray.700"
            p={2}
            borderRadius="md"
            boxShadow="sm"
            mb={2}
            cursor="pointer"
            _hover={{ bg: "gray.600", boxShadow: "md", borderColor: "blue.400" }}
            border="1px solid"
            borderColor="whiteAlpha.100"
            transition="all 0.5s ease-in-out"
            onMouseEnter={() => setIsHovered(true)}
            onMouseLeave={() => setIsHovered(false)}
            role="group"
        >
            {card.cover && (
                <Image
                    src={card.cover}
                    alt={card.title}
                    borderRadius="sm"
                    mb={2}
                    objectFit="cover"
                    w="full"
                    h="120px"
                />
            )}
            <Flex align="start" gap={2}>
                {isHovered && !card.completed && (
                    <Icon
                        as={FiCircle}
                        color="gray.400"
                        boxSize={4}
                        mt={0.5}
                        animation="fadeIn 0.3s ease-out 0.1s both"
                        onClick={handleToggleComplete}
                        cursor="pointer"
                        _hover={{ color: "green.400" }}
                    />
                )}
                {card.completed && (
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
                    fontSize="sm"
                    color="white"
                    flex={1}
                    textDecoration={card.completed ? "line-through" : "none"}
                    opacity={card.completed ? 0.7 : 1}
                    animation={isHovered ? "slideIn 0.4s ease-out 0.2s both" : "none"}
                >
                    {card.title}
                </Text>
                {isHovered && (
                    <Icon
                        as={FiEdit2}
                        color="gray.400"
                        boxSize={3.5}
                        _hover={{ color: "white" }}
                        mt={0.5}
                        animation="fadeIn 0.3s ease-out 0.3s both"
                    />
                )}
            </Flex>
            {
                //card.completed && (
                //    <HStack gap={2} mt={1}>
                //        <Badge
                //            colorPalette="green"
                //            variant="solid"
                //            size="xs"
                //            animation="pulse 2s infinite 0.5s"
                //        >
                //            <Icon
                //                as={FiCheckSquare}
                //                mr={1}
                //                animation="checkBounce 0.6s ease-out"
                //            />
                //            Done
                //        </Badge>
                //    </HStack>
                //)
            }
        </Box>
    );
}
