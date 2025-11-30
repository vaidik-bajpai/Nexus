import { Box, Button, Field, Flex, Heading, Icon, Input, Popover, Portal } from "@chakra-ui/react";
import { FiX } from "react-icons/fi";

const AddCheckList = () => {
    return (
        <Portal>
            <Popover.Positioner>
                <Popover.Content width="auto" p={0} borderRadius="md" boxShadow="lg" zIndex={1600} onMouseDown={(e) => e.stopPropagation()}>
                    <Popover.Body p={0}>
                        <Box w={"xs"} color="gray.700" px={4}>
                            <Flex align="center" justify="space-between" mb={4} px={2} mt={2}>
                                <Box w={8} /> {/* Spacer for centering */}
                                <Heading size="sm" fontWeight="semibold" flex={1} textAlign="center">Add checklist</Heading>
                                <Popover.CloseTrigger as={"div"}>
                                    <Button
                                        size="xs"
                                        variant="ghost"
                                        color="gray.500"
                                        _hover={{ color: "gray.800" }}
                                        p={0}
                                        minW={8}
                                    >
                                        <Icon as={FiX} boxSize={4} />
                                    </Button>
                                </Popover.CloseTrigger>
                            </Flex>

                            <Field.Root pb={5}>
                                <Field.Label>Title</Field.Label>
                                <Input placeholder="Checklist name" size={"xs"} />
                            </Field.Root>
                            <Button size={"2xs"} px={6} py={4} background={"blue.500"} _hover={{ background: "blue.600" }} fontSize={"sm"}>
                                Add
                            </Button>
                        </Box>
                    </Popover.Body>
                </Popover.Content>
            </Popover.Positioner>
        </Portal>
    );
};

export default AddCheckList;