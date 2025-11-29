import { Box, Button, Flex, Heading, Icon, Input, Popover, Portal } from "@chakra-ui/react";
import { FiX } from "react-icons/fi";

const ChangeMembers = () => {
    return (
        <Portal>
            <Popover.Positioner>
                <Popover.Content width="auto" p={0} borderRadius="md" boxShadow="lg" zIndex={1600} onMouseDown={(e) => e.stopPropagation()}>
                    <Popover.Body p={0}>
                        <Box w={"xs"} color="gray.700" >
                            <Flex align="center" justify="space-between" mb={4} px={2} mt={2}>
                                <Box w={8} /> {/* Spacer for centering */}
                                <Heading size="sm" fontWeight="semibold" flex={1} textAlign="center">Change members</Heading>
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
                            </Flex>

                            <Box mb={4} mx={3}>
                                <Input placeholder="Search members" />
                            </Box>
                        </Box>
                    </Popover.Body>
                </Popover.Content>
            </Popover.Positioner>
        </Portal>
    );
};

export default ChangeMembers;