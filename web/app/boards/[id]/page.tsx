"use client"

import { useEffect, useState, use } from "react";
import { getBoard } from "@/lib/services/board";
import { BoardDetail } from "@/lib/types/board.types";
import BoardLayout from "@/components/Board/BoardLayout";
import BoardList from "@/components/Board/BoardList";
import CreateList from "@/components/Board/CreateList";
import { Flex, Spinner, Center, Text, Box, Icon } from "@chakra-ui/react";
import { closestCenter, closestCorners, DndContext, DragEndEvent, DragOverEvent, DragOverlay, DragStartEvent, KeyboardSensor, PointerSensor, pointerWithin, rectIntersection, TouchSensor, useSensor, useSensors } from "@dnd-kit/core";
import { arrayMove, sortableKeyboardCoordinates } from "@dnd-kit/sortable";
import { Card } from "@/lib/types/cards.types";
import BoardCardOverlay from "@/components/Board/BoardCardOverlay";

interface PageProps {
    params: Promise<{ id: string }>;
}

export default function BoardPage({ params }: PageProps) {
    const { id } = use(params);
    const [board, setBoard] = useState<BoardDetail | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const [activeCard, setIsActiveCard] = useState<Card | null>(null);
    const fetchBoard = async () => {
        try {
            setLoading(true);
            const response = await getBoard(id);
            console.log(response)
            setBoard(response.data.board);
        } catch (err) {
            console.error("Failed to fetch board:", err);
            setError("Failed to load board details.");
        } finally {
            setLoading(false);
        }
    };
    useEffect(() => {
        if (id) {
            fetchBoard();
        }
    }, [id]);

    function handleDragStart(event: DragStartEvent) {
        const cardId = event.active.id;
        const card = board?.lists?.flatMap((list) => list.cards).find((card) => card.id === cardId);
        if (card) {
            setIsActiveCard(card);
        }

    }

    function handleDragEnd(event: DragEndEvent) {
        const { active, over } = event;
        const activeCardId = active.id;
        const overId = over?.id;

        if (!overId) return;

        const findColumn = (id: string | number) => {
            return board?.lists.find((list) => list.id === id) ||
                board?.lists.find((list) => list.cards.some((card) => card.id === id));
        };

        const sourceColumn = findColumn(activeCardId);
        const destinationColumn = findColumn(overId);

        if (!sourceColumn || !destinationColumn) {
            return;
        }

        if (sourceColumn.id === destinationColumn.id) {
            const activeIndex = sourceColumn.cards.findIndex((card) => card.id === activeCardId);
            let overIndex = destinationColumn.cards.findIndex((card) => card.id === overId);

            // If dropping on the list container itself, usually means append to end or it's empty
            if (overId === destinationColumn.id) {
                overIndex = destinationColumn.cards.length - 1;
                // If list is empty, overIndex is -1, which is fine, we just want to ensure it's in the list.
                // But arrayMove needs valid indices.
                // If we are just reordering in same list, and we dropped on the list container...
                // It's ambiguous where to put it. But usually Sortable handles this via closestCorners.
                // If we are here, it means we are in the same list.
                // If overId is the list ID, it might be because we dragged to an empty space in the list.
                // Let's assume we don't need to do anything if indices are same.
            }

            if (activeIndex !== overIndex && overIndex !== -1) {
                setBoard((prevBoard) => {
                    if (!prevBoard?.lists) {
                        return prevBoard;
                    }
                    const newLists = [...prevBoard?.lists];
                    const list = newLists.find((list) => list.id === sourceColumn.id);
                    if (list) {
                        const cards = [...list.cards];
                        // Use arrayMove for reordering
                        list.cards = arrayMove(cards, activeIndex, overIndex);
                        return {
                            ...prevBoard,
                            lists: newLists
                        };
                    }
                    return prevBoard;
                })
            }
        }
    }

    function handleDragOver(event: DragOverEvent) {
        const { active, over } = event;
        if (!over) return;

        const activeCardId = active.id;
        const overId = over.id;

        // Find the containers
        const findColumn = (id: string | number) => {
            return board?.lists.find((list) => list.id === id) ||
                board?.lists.find((list) => list.cards.some((card) => card.id === id));
        };

        const sourceColumn = findColumn(activeCardId);
        const targetColumn = findColumn(overId);

        if (!sourceColumn || !targetColumn) return;

        if (sourceColumn.id !== targetColumn.id) {
            console.log(`swap columns ${sourceColumn.name} and ${targetColumn.name}`)
            setBoard((prevBoard) => {
                if (!prevBoard?.lists) {
                    return prevBoard;
                }
                const newLists = [...prevBoard.lists];

                // Find source and target lists in the new state
                const sourceListIndex = newLists.findIndex(l => l.id === sourceColumn.id);
                const targetListIndex = newLists.findIndex(l => l.id === targetColumn.id);

                if (sourceListIndex === -1 || targetListIndex === -1) return prevBoard;

                const sourceList = newLists[sourceListIndex];
                const targetList = newLists[targetListIndex];

                const cardIndex = sourceList.cards.findIndex((card) => card.id === activeCardId);
                if (cardIndex === -1) return prevBoard;

                const [cardToMove] = sourceList.cards.splice(cardIndex, 1);

                // If over a card, insert at that position. If over a list, insert at end (or 0 if empty).
                // But handleDragOver is mostly for moving between containers.
                // If we are over a list container directly, we usually append.
                // However, dnd-kit's sortable strategy handles index calculation better in dragEnd usually,
                // but for cross-container dragOver, we need to move the item into the new container's items array
                // so the sortable context updates.

                let insertIndex = targetList.cards.length;
                if (overId !== targetColumn.id) {
                    const overCardIndex = targetList.cards.findIndex((card) => card.id === overId);
                    if (overCardIndex !== -1) {
                        insertIndex = overCardIndex;
                        // Adjust index based on direction? For simple swap, just inserting at overIndex is usually fine for Sortable.
                        // But we need to be careful not to flicker.
                        // Let's just insert at the overIndex.
                    }
                }

                targetList.cards.splice(insertIndex, 0, cardToMove);

                return {
                    ...prevBoard,
                    lists: newLists
                }
            })
        }
    }

    const sensors = useSensors(
        useSensor(PointerSensor),
        useSensor(TouchSensor),
        useSensor(KeyboardSensor, {
            coordinateGetter: sortableKeyboardCoordinates,
        })
    );

    if (loading) {
        return (
            <Center h="100vh" w="100vw" bg="gray.900">
                <Spinner size="xl" color="blue.500" />
            </Center>
        );
    }

    if (error || !board) {
        return (
            <Center h="100vh" w="100vw" bg="gray.900" color="white">
                <Text>{error || "Board not found"}</Text>
            </Center>
        );
    }

    return (
        <BoardLayout background={board.background} title={board.name}>
            <DndContext
                sensors={sensors}
                collisionDetection={closestCorners}
                onDragStart={handleDragStart}
                onDragOver={handleDragOver}
                onDragEnd={handleDragEnd}
            >
                <Flex h="full" align="flex-start">

                    {board.lists && board.lists.map((list) => (
                        <BoardList key={list.id} list={list} boardId={board.id} onCardCreated={fetchBoard} />
                    ))}
                    <DragOverlay>{activeCard && <BoardCardOverlay card={activeCard} listId={board.lists?.find((list) => list.cards.includes(activeCard))?.id || ""} boardId={board.id} onUpdate={fetchBoard} />}</DragOverlay>
                    <CreateList boardId={board.id} onListCreated={fetchBoard} />
                </Flex>
            </DndContext>
        </BoardLayout>
    );
}
