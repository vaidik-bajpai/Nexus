import { Box, Button, Checkbox, Flex, Heading, Icon, Input, Popover, Portal, Text, Grid, Theme } from "@chakra-ui/react"
import { FiEdit2, FiX, FiChevronLeft } from "react-icons/fi"
import { useState } from "react"
import { createLabel } from "@/lib/services/labels"
import { toaster } from "@/components/ui/toaster"

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
}

const LabelPortal = ({ boardId }: LabelPortalProps) => {
    const [isCreating, setIsCreating] = useState(false);
    const [title, setTitle] = useState("");
    const [selectedColor, setSelectedColor] = useState("");

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
            await createLabel({ boardID: boardId, name: title, color: selectedColor });
            toaster.create({
                title: "Label created successfully",
                type: "success",
            });
            setIsCreating(false);
            setTitle("");
            setSelectedColor("");
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
                <Popover.Content width="300px" p={0} borderRadius="md" boxShadow="lg" zIndex={1600} onMouseDown={(e) => e.stopPropagation()} bg="gray.800" border="1px solid" borderColor="gray.700">
                    <Popover.Body p={0}>
                        <Box color="gray.300" px={4} pb={4} pt={2}>
                            <Flex align="center" justify="space-between" mb={4} px={0} borderBottom="1px solid" borderColor="gray.700" pb={2}>
                                {isCreating ? (
                                    <Button
                                        size="xs"
                                        variant="ghost"
                                        color="gray.400"
                                        _hover={{ color: "white", bg: "gray.700" }}
                                        p={0}
                                        minW={8}
                                        onClick={() => setIsCreating(false)}
                                    >
                                        <Icon as={FiChevronLeft} boxSize={4} />
                                    </Button>
                                ) : (
                                    <Box w={8} />
                                )}
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

                            {isCreating ? (
                                <Flex direction="column" gap={3}>
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
                                        onClick={handleCreateLabel}
                                    >
                                        Create
                                    </Button>
                                </Flex>
                            ) : (
                                <>
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
                                        {colors.slice(0, 4).map((color) => (
                                            <LabelItem key={color.normal} color={color.normal} hoverColor={color.hover} />
                                        ))}
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
                                </>
                            )}
                        </Box>
                    </Popover.Body>
                </Popover.Content>
            </Popover.Positioner>
        </Portal>
    )
}

const LabelItem = ({ color, hoverColor }: { color: string, hoverColor: string }) => {
    return (
        <Flex w={"full"} align="center" justify="space-between" gap={1}>
            <Checkbox.Root variant={"subtle"} w={"full"} size={"sm"}>
                <Checkbox.HiddenInput />
                <Checkbox.Control cursor={"pointer"} mr={1} borderColor="gray.600" _checked={{ bg: "blue.500", borderColor: "blue.500" }} />
                <Checkbox.Label asChild>
                    <Box h={"32px"} w={"full"} bg={color} borderRadius={"sm"} cursor={"pointer"} _hover={{ bg: hoverColor }}></Box>
                </Checkbox.Label>
            </Checkbox.Root>
            <Button as={FiEdit2} size="xs" variant="ghost" color="gray.400" _hover={{ backgroundColor: "gray.700", color: "white" }} />
        </Flex >
    )
}

export default LabelPortal;