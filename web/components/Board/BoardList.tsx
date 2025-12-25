import { Box, Flex, Text, IconButton, Icon } from "@chakra-ui/react";
import { List } from "@/lib/types/list.types";
import BoardCard from "./BoardCard";
import CreateCard from "./CreateCard";
import { FiMoreHorizontal } from "react-icons/fi";
import { SortableContext, useSortable, verticalListSortingStrategy } from "@dnd-kit/sortable";
import { CSS } from "@dnd-kit/utilities";

import { Card } from "@/lib/types/cards.types";

interface BoardListProps {
    list: List;
    cards: Card[];
    boardId: string;
    onCardCreated: () => void;
    onCardClick: (card: Card) => void;
}

export default function BoardList({ list, cards, boardId, onCardCreated, onCardClick }: BoardListProps) {
    const {
        setNodeRef,
        attributes,
        listeners,
        transform,
        transition,
        isDragging,
    } = useSortable({
        id: list.id,
        data: {
            type: "Column",
            list,
        },
    });

    const style = {
        transition,
        transform: CSS.Translate.toString(transform),
        opacity: isDragging ? 0.5 : 1,
    };

    return (
        <Box
            ref={setNodeRef}
            style={style}
            {...attributes}
            {...listeners}
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
        >
            <Flex align="center" justify="space-between" mb={2} px={2} cursor="grab">
                <Text fontWeight="bold" fontSize="sm" color="white">
                    {list.name}
                </Text>
                <IconButton
                    aria-label="List actions"
                    variant="ghost"
                    size="xs"
                    color="gray.400"
                    _hover={{ bg: "gray.700", color: "white" }}
                >
                    <Icon as={FiMoreHorizontal} />
                </IconButton>
            </Flex>

            <SortableContext items={cards.map((card) => card.id)} strategy={verticalListSortingStrategy}>
                <Box flex={1} overflowY="auto" px={1} className="custom-scrollbar">
                    {cards.map((card) => (
                        <BoardCard key={`${card.id}-${card.coverSize}`} card={card} listId={list.id} boardId={boardId} onUpdate={onCardCreated} onClick={() => onCardClick(card)} />
                    ))}
                </Box>
            </SortableContext>

            <CreateCard listId={list.id} boardId={boardId} onCardCreated={onCardCreated} cards={cards} />
        </Box>
    );
}
