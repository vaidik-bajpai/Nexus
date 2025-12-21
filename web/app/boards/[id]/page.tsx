"use client"

import { useEffect, useState, use, useCallback } from "react";
import { getBoardDetails, updateList, getCardAndLists } from "@/lib/services/board";
import BoardLayout from "@/components/Board/BoardLayout";
import BoardList from "@/components/Board/BoardList";
import CreateList from "@/components/Board/CreateList";
import { Flex, Spinner, Center, Text } from "@chakra-ui/react";
import { closestCorners, DndContext, DragEndEvent, DragOverEvent, DragOverlay, DragStartEvent, PointerSensor, TouchSensor, useSensor, useSensors, UniqueIdentifier } from "@dnd-kit/core";
import { arrayMove, horizontalListSortingStrategy, SortableContext } from "@dnd-kit/sortable";
import { Card } from "@/lib/types/cards.types";
import { List } from "@/lib/types/list.types";
import BoardCardOverlay from "@/components/Board/BoardCardOverlay";
import BoardListOverlay from "@/components/Board/BoardListOverlay";
import { updateCard } from "@/lib/services/cards";
import CardModal from "@/components/CardModal/CardModal";
import { useBoardStore } from "@/lib/store/board";

interface PageProps {
    params: Promise<{ id: string }>;
}

export default function BoardPage({ params }: PageProps) {
    const { id } = use(params);
    const { metadata, lists, cards, setMetadata, setCardsAndLists } = useBoardStore();
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const [activeCard, setIsActiveCard] = useState<Card | null>(null);
    const [activeList, setIsActiveList] = useState<List | null>(null);
    const [selectedCard, setSelectedCard] = useState<Card | null>(null);

    const fetchCompleteBoard = useCallback(async () => {
        try {
            const [boardDetails, cardAndLists] = await Promise.all([getBoardDetails(id), getCardAndLists(id)]);
            const { boardID, cards, lists } = cardAndLists.data.board;
            const { board: metadata } = boardDetails.data;

            setMetadata(metadata);
            setCardsAndLists(boardID, cards || [], lists || []);
        } catch (err) {
            console.error("Failed to fetch board:", err);
            setError("Failed to load board details.");
        } finally {
            setLoading(false);
        }
    }, [id, setMetadata, setCardsAndLists]);

    useEffect(() => {
        if (id) {
            fetchCompleteBoard();
        }
    }, [id, fetchCompleteBoard]);

    function handleDragStart(event: DragStartEvent) {
        if (event.active.data.current?.type === "Column") {
            setIsActiveList(event.active.data.current.list);
            return;
        }

        const cardId = event.active.id;
        const card = cards.find((card) => card.id === cardId);
        if (card) {
            setIsActiveCard(card);
        }
    }

    const handleDragOver = (event: DragOverEvent) => {
        const { active, over } = event;
        if (!over) return;

        const activeId = active.id;
        const overId = over.id;

        if (active.data.current?.type === "Column") {
            return;
        }

        const findContainer = (id: UniqueIdentifier) => {
            if (lists.find((list) => list.id === id)) {
                return id as string;
            }
            return cards.find((card) => card.id === id)?.list_id;
        };

        const activeContainer = findContainer(activeId);
        const overContainer = findContainer(overId);

        if (!activeContainer || !overContainer || activeContainer === overContainer) {
            return;
        }

        const newCards = cards.map(c => {
            if (c.id === activeId) {
                return { ...c, list_id: overContainer as string };
            }
            return c;
        });

        setCardsAndLists(id, newCards, lists);
    }

    const handleDragEnd = async (event: DragEndEvent) => {
        const { active, over } = event;

        if (!over) {
            setIsActiveCard(null);
            setIsActiveList(null);
            return;
        }

        if (active.data.current?.type === "Column") {
            if (active.id !== over.id) {
                const oldIndex = lists.findIndex(l => l.id === active.id);
                const newIndex = lists.findIndex(l => l.id === over.id);

                if (oldIndex !== -1 && newIndex !== -1) {
                    const newLists = arrayMove(lists, oldIndex, newIndex);
                    setCardsAndLists(id, cards, newLists);

                    let newPos = 65536.0;
                    if (newLists.length > 1) {
                        if (newIndex === 0) {
                            newPos = newLists[1].position / 2;
                        } else if (newIndex === newLists.length - 1) {
                            newPos = newLists[newIndex - 1].position + 65536.0;
                        } else {
                            newPos = (newLists[newIndex - 1].position + newLists[newIndex + 1].position) / 2;
                        }
                    }

                    try {
                        await updateList({
                            listID: active.id as string,
                            boardID: id,
                            position: newPos
                        });
                    } catch (error) {
                        console.error("Failed to update list position", error);
                    }
                }
            }
            setIsActiveList(null);
            return;
        }

        const activeId = active.id;
        const overId = over.id;

        const findContainer = (id: UniqueIdentifier) => {
            if (lists.find((list) => list.id === id)) {
                return id as string;
            }
            return cards.find((card) => card.id === id)?.list_id;
        };

        const activeContainer = findContainer(activeId);
        const overContainer = findContainer(overId);

        if (activeContainer && overContainer) {
            const overListCards = cards.filter(c => c.list_id === overContainer);

            const activeIndex = cards.findIndex(c => c.id === activeId);
            const overIndex = overListCards.findIndex(c => c.id === overId);

            let newIndex;
            if (lists.find(l => l.id === overId)) {
                newIndex = overListCards.length;
            } else {
                const isBelowOverItem =
                    over &&
                    active.rect.current.translated &&
                    active.rect.current.translated.top >
                    over.rect.top + over.rect.height;

                const modifier = isBelowOverItem ? 1 : 0;
                newIndex = overIndex >= 0 ? overIndex + modifier : overListCards.length;
            }

            const activeCardInListIndex = overListCards.findIndex(c => c.id === activeId);

            if (activeCardInListIndex !== -1) {
                let newOrderedCards = arrayMove(overListCards, activeCardInListIndex, newIndex);

                const cardsWithoutActive = overListCards.filter(c => c.id !== activeId);

                if (lists.find(l => l.id === overId)) {
                    newIndex = overListCards.length - 1;
                } else {
                    const isBelowOverItem =
                        over &&
                        active.rect.current.translated &&
                        active.rect.current.translated.top >
                        over.rect.top + over.rect.height;

                    const modifier = isBelowOverItem ? 1 : 0;
                    newIndex = overIndex >= 0 ? overIndex + modifier : overListCards.length - 1;
                }

                if (newIndex < 0) newIndex = 0;
                if (newIndex >= overListCards.length) newIndex = overListCards.length - 1;

                if (activeCardInListIndex !== newIndex) {
                    newOrderedCards = arrayMove(overListCards, activeCardInListIndex, newIndex);

                    const otherCards = cards.filter(c => c.list_id !== overContainer);
                    const finalCards = [...otherCards, ...newOrderedCards];

                    setCardsAndLists(id, finalCards, lists);

                    let newPos = 65536.0;
                    if (newOrderedCards.length > 1) {
                        if (newIndex === 0) {
                            newPos = newOrderedCards[1].position / 2;
                        } else if (newIndex === newOrderedCards.length - 1) {
                            newPos = newOrderedCards[newIndex - 1].position + 65536.0;
                        } else {
                            newPos = (newOrderedCards[newIndex - 1].position + newOrderedCards[newIndex + 1].position) / 2;
                        }
                    } else {
                        newPos = 65536.0;
                    }

                    try {
                        await updateCard({
                            cardID: activeId as string,
                            listID: overContainer as string,
                            boardID: id,
                            position: newPos
                        });
                    } catch (error) {
                        console.error("Failed to update card position", error);
                    }
                } else {
                    let newPos = 65536.0;
                    if (overListCards.length > 1) {
                        if (newIndex === 0) {
                            newPos = overListCards[1].position / 2;
                        } else if (newIndex === overListCards.length - 1) {
                            newPos = overListCards[newIndex - 1].position + 65536.0;
                        } else {
                            newPos = (overListCards[newIndex - 1].position + overListCards[newIndex + 1].position) / 2;
                        }
                    }

                    try {
                        await updateCard({
                            cardID: activeId as string,
                            listID: overContainer as string,
                            boardID: id,
                            position: newPos
                        });
                    } catch (error) {
                        console.error("Failed to update card list/position", error);
                    }
                }
            }
        }

        setIsActiveCard(null);
        setIsActiveList(null);
    }

    const sensors = useSensors(
        useSensor(PointerSensor, {
            activationConstraint: {
                distance: 5,
            },
        }),
        useSensor(TouchSensor, {
            activationConstraint: {
                delay: 250,
                tolerance: 5,
            },
        })
    );

    if (loading) {
        return (
            <Center h="100vh" w="100vw" bg="gray.900">
                <Spinner size="xl" color="blue.500" />
            </Center>
        );
    }

    if (error || !metadata) {
        return (
            <Center h="100vh" w="100vw" bg="gray.900" color="white">
                <Text>{error || "Board not found"}</Text>
            </Center>
        );
    }

    return (
        <BoardLayout background={metadata.background} title={metadata.name}>
            <DndContext
                sensors={sensors}
                collisionDetection={closestCorners}
                onDragStart={handleDragStart}
                onDragOver={handleDragOver}
                onDragEnd={handleDragEnd}
            >
                <Flex h="full" align="flex-start">
                    <SortableContext items={lists.map(l => l.id)} strategy={horizontalListSortingStrategy}>
                        {lists.map((list) => (
                            <BoardList
                                key={list.id}
                                list={list}
                                cards={cards.filter(c => c.list_id === list.id)}
                                boardId={id}
                                onCardCreated={fetchCompleteBoard}
                                onCardClick={(card) => setSelectedCard(card)}
                            />
                        ))}
                    </SortableContext>
                    <DragOverlay>
                        {activeCard && <BoardCardOverlay card={activeCard} listId={activeCard.list_id} boardId={id} onUpdate={fetchCompleteBoard} />}
                        {activeList && <BoardListOverlay list={activeList} />}
                    </DragOverlay>
                    <CreateList boardId={id} onListCreated={fetchCompleteBoard} lists={lists} />
                </Flex>
            </DndContext>
            {selectedCard && (
                <CardModal
                    isOpen={!!selectedCard}
                    onClose={() => setSelectedCard(null)}
                    card={selectedCard}
                    listName={lists.find(l => l.id === selectedCard.list_id)?.name || ""}
                    boardId={id}
                    listId={selectedCard.list_id}
                    onUpdate={() => {
                        fetchCompleteBoard();
                    }}
                />
            )}
        </BoardLayout>
    );
}

