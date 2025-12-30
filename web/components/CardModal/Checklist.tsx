import { Box, Button, Checkbox, Flex, Icon, Input, Progress, Text } from "@chakra-ui/react";
import { useState } from "react";
import { FiCheckSquare } from "react-icons/fi";
import { Checklist as ChecklistType, ChecklistItem } from "@/lib/types/cards.types";

interface ChecklistProps {
    checklist: ChecklistType;
    onDelete: (id: string) => void;
    onUpdate: (id: string, data: Partial<ChecklistType>) => void;
}

const Checklist = ({ checklist, onDelete, onUpdate }: ChecklistProps) => {
    const [newItemTitle, setNewItemTitle] = useState("");
    const [isAddingItem, setIsAddingItem] = useState(false);

    const items = checklist.checkItems || [];
    const completedCount = items.filter(i => i.completed).length;
    const totalCount = items.length;
    const progress = totalCount === 0 ? 0 : Math.round((completedCount / totalCount) * 100);

    const handleAddItem = () => {
        if (newItemTitle.trim()) {
            const newItem: ChecklistItem = {
                id: crypto.randomUUID(),
                name: newItemTitle,
                completed: false,
                position: items.length + 1
            };
            onUpdate(checklist.id, {
                checkItems: [...items, newItem]
            });
            setNewItemTitle("");
        }
    };

    const handleToggleItem = (itemId: string) => {
        const newItems = items.map(item =>
            item.id === itemId ? { ...item, completed: !item.completed } : item
        );
        onUpdate(checklist.id, { checkItems: newItems });
    };

    const handleDeleteItem = (itemId: string) => {
        const newItems = items.filter(item => item.id !== itemId);
        onUpdate(checklist.id, { checkItems: newItems });
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
                <Button
                    size="xs"
                    variant="subtle"
                    bg="gray.700"
                    color="gray.300"
                    _hover={{ bg: "gray.600" }}
                    onClick={() => onDelete(checklist.id)}
                >
                    Delete
                </Button>
            </Flex>

            <Flex align="center" gap={3} mb={3}>
                <Text fontSize="xs" color="gray.500" w="20px" textAlign="right">
                    {progress}%
                </Text>
                <Progress.Root value={progress} size="sm" flex={1} colorPalette="blue">
                    <Progress.Track bg="gray.700">
                        <Progress.Range />
                    </Progress.Track>
                </Progress.Root>
            </Flex>

            <Box pl={8}>
                {items.map((item) => (
                    <Flex key={item.id} align="center" gap={2} mb={1.5}>
                        <Checkbox.Root
                            checked={item.completed}
                            onCheckedChange={() => handleToggleItem(item.id)}
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
                            onClick={() => handleToggleItem(item.id)}
                            cursor="pointer"
                        >
                            {item.name}
                        </Text>
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
        </Box>
    );
};

export default Checklist;
