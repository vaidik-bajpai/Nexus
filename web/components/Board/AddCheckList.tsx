import { Box, Button, Flex, Heading, Icon, Input, Popover, Portal, Text } from "@chakra-ui/react";
import { useRef, useState } from "react";
import { FiX } from "react-icons/fi";
import { addChecklistToCard } from "@/lib/services/cards";
import { toaster } from "../ui/toaster";

interface AddCheckListProps {
    cardID: string;
    listID: string;
    boardID: string;
    onAdded?: () => void;
}

const AddCheckList = ({ cardID, listID, boardID, onAdded }: AddCheckListProps) => {
    const [name, setName] = useState("Checklist");

    const handleAddChecklist = async () => {
        try {
            await addChecklistToCard({
                cardID,
                listID,
                boardID,
                name,
            })
            onAdded?.();
        } catch (error) {
            toaster.create({
                title: "Failed to add checklist",
                type: "error",
            })
        }
    }

    return (
        <Popover.Positioner>
            <Popover.Content
                width="300px"
                maxH="calc(100vh - 5rem)"
                overflowY="auto"
                p={0}
                borderRadius="md"
                boxShadow="lg"
                onMouseDown={(e) => e.stopPropagation()}
                onClick={(e) => e.stopPropagation()}
                bg="gray.800"
                border="1px solid"
                borderColor="gray.700"
            >
                <Box color="gray.300" px={4} pb={4} pt={2}>
                    <Flex align="center" justify="space-between" mb={4} px={0} borderBottom="1px solid" borderColor="gray.700" pb={2}>
                        <Box w={8} />
                        <Heading size="sm" fontWeight="semibold" flex={1} textAlign="center" color="gray.300">
                            Add checklist
                        </Heading>
                        <Popover.CloseTrigger asChild>
                            <Button
                                size="xs"
                                variant="ghost"
                                color="gray.400"
                                _hover={{ color: "white", bg: "gray.700" }}
                                p={0}
                                minW={8}
                                onClick={(e) => {
                                    e.stopPropagation();
                                }}
                            >
                                <Icon as={FiX} boxSize={4} />
                            </Button>
                        </Popover.CloseTrigger>
                    </Flex>

                    <Box>
                        <Text fontSize="xs" fontWeight="bold" color="gray.500" mb={1.5}>Title</Text>
                        <Input
                            value={name}
                            onChange={(e) => setName(e.target.value)}
                            bg="gray.900"
                            border="1px solid"
                            borderColor="gray.700"
                            _focus={{ borderColor: "blue.500", outline: "none" }}
                            size="sm"
                            color="white"
                            onKeyDown={(e) => {
                                if (e.key === "Enter") {
                                    handleAddChecklist();
                                }
                            }}
                        />
                    </Box>

                    <Popover.CloseTrigger asChild>
                        <Button
                            size="sm"
                            colorPalette="blue"
                            w="fit-content"
                            mt={3}
                            onClick={(e) => {
                                handleAddChecklist();
                            }}
                        >
                            Add
                        </Button>
                    </Popover.CloseTrigger>
                </Box>
            </Popover.Content>
        </Popover.Positioner>
    );
};

export default AddCheckList;