import { Box, Flex, Text, IconButton, Icon } from "@chakra-ui/react";
import { List } from "@/lib/types/list.types";
import BoardCard from "./BoardCard";
import { FiMoreHorizontal } from "react-icons/fi";

import { Card } from "@/lib/types/cards.types";

interface BoardListOverlayProps {
    list: List;
    cards: Card[];
}

export default function BoardListOverlay({ list, cards }: BoardListOverlayProps) {
    return (
        <Box
            w="272px"
            minW="272px"
            backdropFilter="blur(4px)"
            borderRadius="xl"
            border="1px solid"
            borderColor="whiteAlpha.100"
            p={2}
            mr={3}
            h="fit-content"
            maxH="full"
            minH="100px"
            display="flex"
            flexDirection="column"
            bg={"blackAlpha.800"}
            cursor="grabbing"
            transform="rotate(4deg)" // Rotation effect
            boxShadow="xl"
        >
            <Flex align="center" justify="space-between" mb={2} px={2}>
                <Text fontWeight="bold" fontSize="sm" color="white">
                    {list.name}
                </Text>
                <IconButton
                    aria-label="List actions"
                    variant="ghost"
                    size="xs"
                    color="gray.400"
                >
                    <Icon as={FiMoreHorizontal} />
                </IconButton>
            </Flex>

            <Box flex={1} overflowY="auto" px={1} className="custom-scrollbar">
                {cards.map((card) => (
                    <BoardCard key={card.id} card={card} listId={list.id} boardId="" onUpdate={() => { }} onClick={() => { }} />
                ))}
            </Box>
        </Box>
    );
}
