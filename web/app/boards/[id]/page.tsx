"use client"

import { useEffect, useState, use } from "react";
import { getBoard } from "@/lib/services/board";
import { BoardDetail } from "@/lib/types/board.types";
import BoardLayout from "@/components/Board/BoardLayout";
import BoardList from "@/components/Board/BoardList";
import CreateList from "@/components/Board/CreateList";
import { Flex, Spinner, Center, Text, Box, Icon } from "@chakra-ui/react";
import { closestCenter, closestCorners, DndContext, DragEndEvent, DragOverEvent, DragOverlay, DragStartEvent, KeyboardSensor, PointerSensor, pointerWithin, rectIntersection, TouchSensor, useSensor, useSensors, UniqueIdentifier } from "@dnd-kit/core";
import { arrayMove, sortableKeyboardCoordinates, horizontalListSortingStrategy, SortableContext } from "@dnd-kit/sortable";
import { Card } from "@/lib/types/cards.types";
import { List } from "@/lib/types/list.types";
import BoardCardOverlay from "@/components/Board/BoardCardOverlay";
import BoardListOverlay from "@/components/Board/BoardListOverlay";
import { updateCard } from "@/lib/services/cards";

interface PageProps {
    params: Promise<{ id: string }>;
}

export default function BoardPage({ params }: PageProps) {
    const { id } = use(params);
    const [board, setBoard] = useState<BoardDetail | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const [activeCard, setIsActiveCard] = useState<Card | null>(null);
    const [activeList, setIsActiveList] = useState<List | null>(null); // Added activeList state
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
        if (event.active.data.current?.type === "Column") {
            setIsActiveList(event.active.data.current.list);
            return;
        }

        const cardId = event.active.id;
        const card = board?.lists?.flatMap((list) => list.cards || []).find((card) => card.id === cardId); // Added || []
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
            // List reordering logic (visual only for now)
            return;
        }

        // Find the containers
        const findContainer = (id: UniqueIdentifier) => {
            if (board?.lists.find((list) => list.id === id)) {
                return id as string;
            }
            return board?.lists.find((list) => list.cards?.some((card) => card.id === id))?.id;
        };

        const activeContainer = findContainer(activeId);
        const overContainer = findContainer(overId);

        if (!activeContainer || !overContainer || activeContainer === overContainer) {
            return;
        }

        setBoard((prev) => {
            if (!prev) return prev;
            const activeItems = prev.lists.find(l => l.id === activeContainer)?.cards || [];
            const overItems = prev.lists.find(l => l.id === overContainer)?.cards || [];

            const activeIndex = activeItems.findIndex(c => c.id === activeId);
            const overIndex = overItems.findIndex(c => c.id === overId);

            let newIndex;
            if (prev.lists.find(l => l.id === overId)) {
                // We're over the container itself
                newIndex = overItems.length + 1;
            } else {
                const isBelowOverItem =
                    over &&
                    active.rect.current.translated &&
                    active.rect.current.translated.top >
                    over.rect.top + over.rect.height;

                const modifier = isBelowOverItem ? 1 : 0;
                newIndex = overIndex >= 0 ? overIndex + modifier : overItems.length + 1;
            }

            return {
                ...prev,
                lists: prev.lists.map(list => {
                    if (list.id === activeContainer) {
                        return { ...list, cards: (list.cards || []).filter(c => c.id !== activeId) };
                    }
                    if (list.id === overContainer) {
                        const newCards = [...(list.cards || [])];
                        const cardToMove = activeItems[activeIndex];
                        if (!cardToMove) return list;

                        newCards.splice(newIndex, 0, cardToMove);
                        return { ...list, cards: newCards };
                    }
                    return list;
                })
            };
        });
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
                setBoard((prev) => {
                    if (!prev) return prev;
                    const activeIndex = prev.lists.findIndex(l => l.id === active.id);
                    const overIndex = prev.lists.findIndex(l => l.id === over.id);
                    return {
                        ...prev,
                        lists: arrayMove(prev.lists, activeIndex, overIndex)
                    };
                });
            }
            setIsActiveList(null);
            return;
        }

        const activeId = active.id;
        const overId = over?.id;

        if (!overId) return;

        const findContainer = (id: UniqueIdentifier) => {
            if (board?.lists.find((list) => list.id === id)) {
                return id as string;
            }
            return board?.lists.find((list) => list.cards?.some((card) => card.id === id))?.id;
        };

        const activeContainer = findContainer(activeId);
        const overContainer = findContainer(overId);

        // Get the original container from dnd-kit's active data
        const originalContainer = active.data.current?.sortable?.containerId;

        if (activeContainer && overContainer) {
            const activeList = board?.lists.find(l => l.id === activeContainer);
            const overList = board?.lists.find(l => l.id === overContainer);

            if (!activeList || !overList) return;

            const activeIndex = activeList.cards.findIndex(c => c.id === activeId);
            const overIndex = overList.cards.findIndex(c => c.id === overId);

            let newIndex;
            if (activeContainer === overContainer) {
                // Same container reordering (or after dragOver moved it)
                newIndex = overIndex;
                if (overId === activeContainer) {
                    newIndex = activeList.cards.length - 1;
                }

                // Check if position changed OR list changed
                // If activeContainer (current) != originalContainer, then list changed!
                const listChanged = originalContainer && activeContainer !== originalContainer;

                if (activeIndex !== newIndex || listChanged) {
                    const newCards = arrayMove(activeList.cards, activeIndex, newIndex);

                    // Optimistic update
                    setBoard((prev) => {
                        if (!prev) return prev;
                        return {
                            ...prev,
                            lists: prev.lists.map(l => {
                                if (l.id === activeContainer) {
                                    return { ...l, cards: newCards };
                                }
                                return l;
                            })
                        };
                    });

                    // Calculate new position
                    let newPos = 65536.0;
                    if (newCards.length > 1) {
                        if (newIndex === 0) {
                            newPos = newCards[1].position / 2;
                        } else if (newIndex === newCards.length - 1) {
                            newPos = newCards[newIndex - 1].position + 65536.0;
                        } else {
                            newPos = (newCards[newIndex - 1].position + newCards[newIndex + 1].position) / 2;
                        }
                    }

                    try {
                        await updateCard({
                            cardID: activeId as string,
                            listID: activeContainer,
                            boardID: board!.id,
                            position: newPos
                        });
                    } catch (error) {
                        console.error("Failed to update card position", error);
                    }
                }
            } else {
                // This block is likely unreachable if handleDragOver updates state, but keeping it safe
                // Cross-container moving
                // The UI is already handled by handleDragOver, we just need to persist the change
                // We need to find the card in the NEW list (overList) because handleDragOver moved it there in state
                const cardInNewList = overList.cards.find(c => c.id === activeId);
                if (!cardInNewList) return; // Should be there due to handleDragOver

                const newIndexInDest = overList.cards.findIndex(c => c.id === activeId);

                // Calculate new position in destination list
                let newPos = 65536.0;
                if (overList.cards.length > 1) {
                    if (newIndexInDest === 0) {
                        newPos = overList.cards[1].position / 2;
                    } else if (newIndexInDest === overList.cards.length - 1) {
                        newPos = overList.cards[newIndexInDest - 1].position + 65536.0;
                    } else {
                        newPos = (overList.cards[newIndexInDest - 1].position + overList.cards[newIndexInDest + 1].position) / 2;
                    }
                }

                try {
                    await updateCard({
                        cardID: activeId as string,
                        listID: overContainer, // New List ID
                        boardID: board!.id,
                        position: newPos
                    });
                    alert("Card updated successfully");
                } catch (error) {
                    alert("Failed to update card list/position");
                    console.error("Failed to update card list/position", error);
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
        }),
        useSensor(KeyboardSensor, {
            coordinateGetter: sortableKeyboardCoordinates,
            keyboardCodes: {
                start: ["Space", "Enter"],
                cancel: ["Escape"],
                end: ["Space", "Enter"],
            },
            onActivation: (event) => {
                // Prevent activation if focused on an input or textarea
                if (document.activeElement?.tagName === "INPUT" || document.activeElement?.tagName === "TEXTAREA") {
                    return;
                }
            }
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
                    <SortableContext items={board.lists.map(l => l.id)} strategy={horizontalListSortingStrategy}>
                        {board.lists && board.lists.map((list) => (
                            <BoardList key={list.id} list={list} boardId={board.id} onCardCreated={fetchBoard} />
                        ))}
                    </SortableContext>
                    <DragOverlay>
                        {activeCard && <BoardCardOverlay card={activeCard} listId={board.lists?.find((list) => list.cards?.includes(activeCard))?.id || ""} boardId={board.id} onUpdate={fetchBoard} />}
                        {activeList && <BoardListOverlay list={activeList} />}
                    </DragOverlay>
                    <CreateList boardId={board.id} onListCreated={fetchBoard} />
                </Flex>
            </DndContext>
        </BoardLayout>
    );
}
