"use client"

import { useEffect, useState, use } from "react";
import { getBoard } from "@/lib/services/board";
import { BoardDetail } from "@/lib/types/board.types";
import BoardLayout from "@/components/Board/BoardLayout";
import BoardList from "@/components/Board/BoardList";
import CreateList from "@/components/Board/CreateList";
import { Flex, Spinner, Center, Text, Box, Icon } from "@chakra-ui/react";
import { closestCenter, DndContext, DragEndEvent, DragOverEvent, DragOverlay, DragStartEvent, KeyboardSensor, PointerSensor, rectIntersection, TouchSensor, useSensor, useSensors } from "@dnd-kit/core";
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
        const overCardId = over?.id;

        const sourceColumn = board?.lists.find((list) => list.cards.some((card) => card.id === activeCardId));
        const destinationColumn = board?.lists.find((list) => list.cards.some((card) => card.id === overCardId));

        if (!sourceColumn || !destinationColumn) {
            return;
        }

        if (sourceColumn.id === destinationColumn.id) {
            const activeIndex = sourceColumn.cards.findIndex((card) => card.id === activeCardId);
            const overIndex = destinationColumn.cards.findIndex((card) => card.id === overCardId);
            if (activeIndex !== overIndex) {
                setBoard((prevBoard) => {
                    if (!prevBoard?.lists) {
                        return prevBoard;
                    }
                    const newLists = [...prevBoard?.lists];
                    const list = newLists.find((list) => list.id === sourceColumn.id);
                    if (list) {
                        const cards = [...list.cards];
                        const [removed] = cards.splice(activeIndex, 1);
                        cards.splice(overIndex, 0, removed);
                        list.cards = cards;
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
        const overCardId = over.id;

        const targetColumn = board?.lists.find((list) => list.cards.some((card) => card.id === overCardId));
        if (targetColumn) {
            const sourceColumn = board?.lists.find((list) => list.cards.some((card) => card.id === activeCardId));
            if (sourceColumn && sourceColumn.id !== targetColumn.id) {
                console.log(`swap columns ${sourceColumn.name} and ${targetColumn.name}`)
                setBoard((prevBoard) => {
                    if (!prevBoard?.lists) {
                        return prevBoard;
                    }
                    const newList = [...prevBoard.lists];
                    let cardToMove: Card | null = null;
                    for (const list of newList) {
                        const cardIndex = list.cards.findIndex((card) => card.id === activeCardId);
                        if (cardIndex !== -1) {
                            cardToMove = list.cards[cardIndex];
                            list.cards.splice(cardIndex, 1);
                            break;
                        }
                    }

                    if (cardToMove) {
                        targetColumn.cards.push(cardToMove);
                    }
                    return {
                        ...prevBoard,
                        lists: newList
                    }
                })
            }
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
                collisionDetection={rectIntersection}
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
