import { Box, Button, Checkbox, Flex, Heading, Icon, Input, Popover, Portal, Text, Grid, Theme } from "@chakra-ui/react"
import { FiEdit2, FiX, FiChevronLeft } from "react-icons/fi"
import { useState } from "react"
import { createLabel, toggleLabel } from "@/lib/services/labels"
import { toaster } from "@/components/ui/toaster"
import { useBoardStore } from "@/lib/store/board"

const colors = [
    { normal: "#116932", hover: "#124a28" },
    { normal: "#845209", hover: "#713f12" },
    { normal: "#92310a", hover: "#6c2710" },
    { normal: "#991919", hover: "#511111" },
    { normal: "#641ba3", hover: "#4a1772" },
    { normal: "#173da6", hover: "#1a3478" },
    { normal: "#0c5c72", hover: "#134152" },
    { normal: "#0c5d56", hover: "#114240" },
    { normal: "#a41752", hover: "#6d0e34" },
    { normal: "#3f3f46", hover: "#27272a" },
]

interface LabelPortalProps {
    boardId?: string;
    cardId?: string;
    listId?: string;
    activeLabels?: any[];
}

const LabelPortal = ({ boardId, cardId, listId, activeLabels }: LabelPortalProps) => {
    const [isCreating, setIsCreating] = useState(false);
    const [title, setTitle] = useState("");
    const [selectedColor, setSelectedColor] = useState("");
    const { metadata, enrichCards, addLabelToBoard } = useBoardStore();
    const labels = metadata?.labels;

    const handleToggleLabel = async (label: any) => {
        if (!cardId || !listId || !boardId) return;

        const isActive = activeLabels?.some(l => l && (l.id === label.id || l.labelID === label.id));
        const newLabels = isActive
            ? activeLabels?.filter(l => l && (l.id !== label.id && l.labelID !== label.id))
            : [...(activeLabels || []), { ...label, labelID: label.id }];

        // Optimistic update
        enrichCards(cardId, { labels: newLabels });

        try {
            await toggleLabel({
                boardID: boardId,
                listID: listId,
                cardID: cardId,
                labelID: label.id,
                type: isActive ? "remove" : "add"
            });
        } catch (error) {
            // Revert on error
            enrichCards(cardId, { labels: activeLabels });
            toaster.create({
                title: "Failed to update label",
                type: "error",
            });
        }
    };

    const handleCreateLabel = async () => {
        if (!boardId) {
            toaster.create({
                title: "Board ID is missing",
                type: "error",
            });
            return;
        }
        if (!title || !selectedColor) {
            toaster.create({
                title: "Please provide a title and select a color",
                type: "error",
            });
            return;
        }

        try {
            const newLabel = await createLabel({ boardID: boardId, name: title, color: selectedColor });
            toaster.create({
                title: "Label created successfully",
                type: "success",
            });
            setIsCreating(false);
            setTitle("");
            setSelectedColor("");

            // Update board state
            addLabelToBoard(newLabel.data);

            // Add the new label to the card
            if (cardId && listId) {
                await handleToggleLabel(newLabel.data);
            }
        } catch (error) {
            toaster.create({
                title: "Failed to create label",
                type: "error",
            });
        }
    };

    return (
        <Portal>
            <Popover.Positioner>
                <Popover.Content width="300px" p={0} borderRadius="md" boxShadow="lg" zIndex={1600} onMouseDown={(e) => e.stopPropagation()} onClick={(e) => e.stopPropagation()} bg="gray.800" border="1px solid" borderColor="gray.700">
                    <Popover.Body p={0}>
                        <Box color="gray.300" px={4} pb={4} pt={2}>
                            <Flex align="center" justify="space-between" mb={4} px={0} borderBottom="1px solid" borderColor="gray.700" pb={2}>
                                <Button
                                    size="xs"
                                    variant="ghost"
                                    color="gray.400"
                                    _hover={{ color: "white", bg: "gray.700" }}
                                    p={0}
                                    minW={8}
                                    display={isCreating ? "inline-flex" : "none"}
                                    onClick={(e) => {
                                        e.stopPropagation();
                                        setIsCreating(false);
                                    }}
                                >
                                    <Icon as={FiChevronLeft} boxSize={4} />
                                </Button>
                                <Box w={8} display={!isCreating ? "block" : "none"} />
                                <Heading size="sm" fontWeight="semibold" flex={1} textAlign="center" color="gray.300">
                                    {isCreating ? "Create label" : "Labels"}
                                </Heading>
                                <Popover.CloseTrigger asChild>
                                    <Button
                                        size="xs"
                                        variant="ghost"
                                        color="gray.400"
                                        _hover={{ color: "white", bg: "gray.700" }}
                                        p={0}
                                        minW={8}
                                        onClick={(e) => e.stopPropagation()}
                                    >
                                        <Icon as={FiX} boxSize={4} />
                                    </Button>
                                </Popover.CloseTrigger>
                            </Flex>

                            <Flex direction="column" gap={3} display={isCreating ? "flex" : "none"}>
                                <Box
                                    h="32px"
                                    w="full"
                                    bg={selectedColor || "gray.700"}
                                    borderRadius="sm"
                                    display="flex"
                                    alignItems="center"
                                    px={3}
                                >
                                    <Text fontSize="sm" fontWeight="medium" color="white">
                                        {title}
                                    </Text>
                                </Box>

                                <Box>
                                    <Text fontSize="xs" fontWeight="bold" color="gray.500" mb={1.5}>Title</Text>
                                    <Input
                                        placeholder="Label title"
                                        value={title}
                                        onChange={(e) => setTitle(e.target.value)}
                                        bg="gray.900"
                                        border="1px solid"
                                        borderColor="gray.700"
                                        _focus={{ borderColor: "blue.500", outline: "none" }}
                                        size="sm"
                                        color="white"
                                    />
                                </Box>

                                <Box>
                                    <Text fontSize="xs" fontWeight="bold" color="gray.500" mb={1.5}>Select a color</Text>
                                    <Grid templateColumns="repeat(5, 1fr)" gap={2}>
                                        {colors.map((color) => (
                                            <Box
                                                key={color.normal}
                                                h="32px"
                                                bg={color.normal}
                                                borderRadius="sm"
                                                cursor="pointer"
                                                _hover={{ bg: color.hover }}
                                                onClick={() => setSelectedColor(color.normal)}
                                                border={selectedColor === color.normal ? "2px solid white" : "none"}
                                            />
                                        ))}
                                    </Grid>
                                </Box>

                                <Button
                                    size="sm"
                                    variant="subtle"
                                    w="full"
                                    mt={2}
                                    onClick={() => setSelectedColor("")}
                                    bg="gray.700"
                                    color="gray.300"
                                    _hover={{ bg: "gray.600" }}
                                >
                                    <Icon as={FiX} mr={1} />
                                    Remove color
                                </Button>

                                <Button
                                    size="sm"
                                    colorPalette="blue"
                                    w="fit-content"
                                    mt={2}
                                    onClick={(e) => {
                                        console.log("create a label");
                                        e.stopPropagation();
                                        handleCreateLabel();
                                    }}
                                >
                                    Create
                                </Button>
                            </Flex>

                            <Box display={!isCreating ? "block" : "none"}>
                                <Input
                                    placeholder="Search labels..."
                                    mb={3}
                                    bg="gray.900"
                                    border="1px solid"
                                    borderColor="gray.700"
                                    _focus={{ borderColor: "blue.500", outline: "none" }}
                                    size="sm"
                                    color="white"
                                />

                                <Flex flexDirection="column" gap={1}>
                                    <Text fontSize="xs" fontWeight="bold" color="gray.500" mb={1}>Labels</Text>
                                    {labels?.map((label) => {
                                        if (!label) return null;
                                        const isChecked = activeLabels?.some(l => l && (l.id === label.id || l.labelID === label.id));
                                        return (
                                            <LabelItem
                                                key={label.id}
                                                color={label.color}
                                                hoverColor={label.color}
                                                text={label.name}
                                                isChecked={isChecked}
                                                onToggle={() => handleToggleLabel(label)}
                                            />
                                        )
                                    })}
                                </Flex>

                                <Button
                                    size="xs"
                                    variant="subtle"
                                    w="full"
                                    mt={4}
                                    fontSize="sm"
                                    fontWeight="semibold"
                                    bg="gray.700"
                                    color="gray.300"
                                    _hover={{ bg: "gray.600" }}
                                    onClick={(e) => {
                                        e.stopPropagation();
                                        e.preventDefault();
                                        setIsCreating(true);
                                    }}
                                >
                                    Create A new label
                                </Button>
                            </Box>
                        </Box>
                    </Popover.Body>
                </Popover.Content>
            </Popover.Positioner>
        </Portal>
    )
}

const LabelItem = ({ color, hoverColor, text, isChecked, onToggle }: { color: string, hoverColor: string, text?: string, isChecked?: boolean, onToggle?: () => void }) => {
    return (
        <Flex w={"full"} align="center" justify="space-between" gap={1} onClick={(e) => e.stopPropagation()}>
            <Checkbox.Root variant={"subtle"} w={"full"} size={"sm"} checked={isChecked} onCheckedChange={() => onToggle?.()}>
                <Checkbox.HiddenInput />
                <Checkbox.Control cursor={"pointer"} mr={1} borderColor="gray.600" _checked={{ bg: "blue.500", borderColor: "blue.500" }} />
                <Checkbox.Label asChild>
                    <Flex h={"32px"} w={"full"} bg={color} borderRadius={"sm"} cursor={"pointer"} _hover={{ bg: hoverColor }} align="center" px={2}>
                        <Text fontSize="xs" color="white" fontWeight="bold">{text}</Text>
                    </Flex>
                </Checkbox.Label>
            </Checkbox.Root>
            <Button as={FiEdit2} size="xs" variant="ghost" color="gray.400" _hover={{ backgroundColor: "gray.700", color: "white" }} onClick={(e) => e.stopPropagation()} />
        </Flex >
    )
}

export default LabelPortal;