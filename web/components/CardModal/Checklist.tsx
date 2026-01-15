import { Box, Button, Checkbox, Flex, Icon, Input, Popover, Portal, Progress, Text } from "@chakra-ui/react";
import { useState } from "react";
import { FiCheckSquare, FiX, FiMoreHorizontal } from "react-icons/fi";
import { Checklist as ChecklistType, ChecklistItem } from "@/lib/types/cards.types";
import * as cardsService from "@/lib/services/cards";
import { toaster } from "@/components/ui/toaster";

interface ChecklistProps {
    checklist: ChecklistType;
    onDelete: (id: string) => void;
    onUpdate: (id: string, data: Partial<ChecklistType>) => void;
    boardId: string;
    listId: string;
    cardId: string;
}

const Checklist = ({ checklist, onDelete, onUpdate, boardId, listId, cardId }: ChecklistProps) => {
    const [newItemTitle, setNewItemTitle] = useState("");
    const [isAddingItem, setIsAddingItem] = useState(false);

    const items = checklist.checkItems || [];
    const completedCount = items.filter(i => i.completed).length;
    const totalCount = items.length;
    const progress = totalCount === 0 ? 0 : Math.round((completedCount / totalCount) * 100);

    const handleAddItem = () => {
        if (newItemTitle.trim()) {
            cardsService.addChecklistItem({
                cardID: cardId,
                checklistID: checklist.id,
                listID: listId,
                boardID: boardId,
                name: newItemTitle
            }).then((res) => {
                const newItem = res.data;
                onUpdate(checklist.id, {
                    checkItems: [...items, newItem]
                });
                setNewItemTitle("");
                // Keep input open for adding more items
            }).catch(() => {
                toaster.create({
                    title: "Failed to add item",
                    type: "error",
                });
            });
        }
    };

    const handleToggleItem = (itemId: string, currentCompleted: boolean) => {
        // Optimistic update
        const newItems = items.map(item =>
            item.id === itemId ? { ...item, completed: !currentCompleted } : item
        );
        onUpdate(checklist.id, { checkItems: newItems });

        cardsService.updateChecklistItem({
            cardID: cardId,
            checklistID: checklist.id,
            listID: listId,
            boardID: boardId,
            itemID: itemId,
            completed: !currentCompleted
        }).catch(() => {
            // Revert on failure
            onUpdate(checklist.id, { checkItems: items });
            toaster.create({
                title: "Failed to update item",
                type: "error",
            });
        });
    };

    const handleDeleteItem = (itemId: string) => {
        // Optimistic update
        const newItems = items.filter(item => item.id !== itemId);
        onUpdate(checklist.id, { checkItems: newItems });

        cardsService.deleteChecklistItem({
            cardID: cardId,
            checklistID: checklist.id,
            listID: listId,
            boardID: boardId,
            itemID: itemId
        }).catch(() => {
            // Revert on failure
            onUpdate(checklist.id, { checkItems: items });
            toaster.create({
                title: "Failed to delete item",
                type: "error",
            });
        });
    };

    return (
        <Box mb={6}>
            <Flex align="center" justify="space-between" mb={2}>
                <Flex align="center" gap={3}>
                    <Icon as={FiCheckSquare} boxSize={5} color="gray.400" />
                    <Text fontWeight="bold" fontSize="md" color="gray.300">
                        {checklist.name}
                    </Text>
                </Flex>
                <Popover.Root positioning={{ placement: "bottom-end" }}>
                    <Popover.Trigger asChild>
                        <Button
                            size="xs"
                            variant="subtle"
                            bg="gray.700"
                            color="gray.300"
                            _hover={{ bg: "gray.600" }}
                        >
                            Delete
                        </Button>
                    </Popover.Trigger>
                    <Portal>
                        <Popover.Positioner>
                            <Popover.Content bg="gray.800" borderColor="gray.700" color="white" w="300px">
                                <Popover.Body p={4}>
                                    <Flex justify="space-between" align="center" mb={2}>
                                        <Text fontWeight="semibold">Delete Checklist?</Text>
                                        <Popover.CloseTrigger asChild>
                                            <Box as="button" color="gray.400" _hover={{ color: "white" }} cursor="pointer">
                                                <Icon as={FiX} />
                                            </Box>
                                        </Popover.CloseTrigger>
                                    </Flex>
                                    <Text fontSize="sm" color="gray.400" mb={4}>
                                        Deleting a checklist is permanent and there is no way to get it back.
                                    </Text>
                                    <Button
                                        size="sm"
                                        colorPalette="red"
                                        w="full"
                                        onClick={() => onDelete(checklist.id)}
                                    >
                                        Delete checklist
                                    </Button>
                                </Popover.Body>
                            </Popover.Content>
                        </Popover.Positioner>
                    </Portal>
                </Popover.Root>
            </Flex>

            <Flex align="center" gap={3} mb={3}>
                <Text fontSize="xs" color="gray.500" textAlign="right">
                    {progress}%
                </Text>
                <Progress.Root value={progress} size="sm" flex={1} colorPalette="blue">
                    <Progress.Track bg="gray.700">
                        <Progress.Range />
                    </Progress.Track>
                </Progress.Root>
            </Flex>

            <Box pl={0}>
                {items.map((item) => (
                    <Flex key={item.id} align="center" gap={2} mb={1.5} role="group">
                        <Checkbox.Root
                            checked={item.completed}
                            onCheckedChange={() => handleToggleItem(item.id, item.completed)}
                            size="md"
                            variant="subtle"
                        >
                            <Checkbox.HiddenInput />
                            <Checkbox.Control
                                borderColor="gray.600"
                                _checked={{ bg: "blue.500", borderColor: "blue.500" }}
                                _hover={{ borderColor: "gray.500" }}
                            />
                        </Checkbox.Root>
                        <Text
                            fontSize="sm"
                            color="gray.300"
                            textDecoration={item.completed ? "line-through" : "none"}
                            opacity={item.completed ? 0.7 : 1}
                            flex={1}
                            onClick={() => handleToggleItem(item.id, item.completed)}
                            cursor="pointer"
                        >
                            {item.name}
                        </Text>
                        <Popover.Root positioning={{ placement: "bottom-end" }}>
                            <Popover.Trigger asChild>
                                <Button
                                    size="xs"
                                    variant="ghost"
                                    color="gray.400"
                                    _hover={{ bg: "gray.700", color: "white" }}
                                >
                                    <Icon as={FiMoreHorizontal} />
                                </Button>
                            </Popover.Trigger>
                            <Portal>
                                <Popover.Positioner>
                                    <Popover.Content bg="gray.800" borderColor="gray.700" color="white" w="200px">
                                        <Popover.Body p={1}>
                                            <Flex justify="space-between" align="center" px={3} py={2} borderBottom="1px solid" borderColor="gray.700" mb={1}>
                                                <Text fontSize="xs" fontWeight="semibold" color="gray.400" textAlign="center" flex={1}>Item actions</Text>
                                                <Popover.CloseTrigger asChild>
                                                    <Box as="button" color="gray.400" _hover={{ color: "white" }} cursor="pointer">
                                                        <Icon as={FiX} boxSize={3} />
                                                    </Box>
                                                </Popover.CloseTrigger>
                                            </Flex>
                                            <Button
                                                size="sm"
                                                variant="ghost"
                                                justifyContent="flex-start"
                                                w="full"
                                                fontWeight="normal"
                                                color="gray.300"
                                                _hover={{ bg: "gray.700" }}
                                                onClick={() => console.log("Convert to card")}
                                            >
                                                Convert to card
                                            </Button>
                                            <Button
                                                size="sm"
                                                variant="ghost"
                                                justifyContent="flex-start"
                                                w="full"
                                                fontWeight="normal"
                                                color="gray.300"
                                                _hover={{ bg: "gray.700" }}
                                                onClick={() => handleDeleteItem(item.id)}
                                            >
                                                Delete
                                            </Button>
                                        </Popover.Body>
                                    </Popover.Content>
                                </Popover.Positioner>
                            </Portal>
                        </Popover.Root>
                    </Flex>
                ))}

                {!isAddingItem ? (
                    <Button
                        size="sm"
                        variant="subtle"
                        bg="gray.700"
                        color="gray.300"
                        _hover={{ bg: "gray.600" }}
                        mt={2}
                        onClick={() => setIsAddingItem(true)}
                    >
                        Add an item
                    </Button>
                ) : (
                    <Box mt={2}>
                        <Input
                            placeholder="Add an item"
                            value={newItemTitle}
                            onChange={(e) => setNewItemTitle(e.target.value)}
                            bg="gray.900"
                            border="1px solid"
                            borderColor="blue.500"
                            _focus={{ borderColor: "blue.500", outline: "none" }}
                            size="sm"
                            color="white"
                            autoFocus
                            mb={2}
                            onKeyDown={(e) => {
                                if (e.key === "Enter") {
                                    handleAddItem();
                                } else if (e.key === "Escape") {
                                    setIsAddingItem(false);
                                }
                            }}
                        />
                        <Flex gap={2} align="center">
                            <Button
                                size="sm"
                                colorPalette="blue"
                                onClick={handleAddItem}
                            >
                                Add
                            </Button>
                            <Button
                                size="sm"
                                variant="ghost"
                                color="gray.400"
                                _hover={{ color: "white" }}
                                onClick={() => setIsAddingItem(false)}
                            >
                                Cancel
                            </Button>
                        </Flex>
                    </Box>
                )}
            </Box>
        </Box >
    );
};

export default Checklist;
