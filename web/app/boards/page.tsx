"use client"

import { Box, Heading, SimpleGrid, Text, Stack, Icon, Popover, Button, Portal } from "@chakra-ui/react";
import BoardsLayout from "@/components/BoardsLayout";
import BoardCard from "@/components/BoardCard";
import { FiClock, FiUsers } from "react-icons/fi";
import CreateBoard from "@/components/CreateBoard";
import { useState } from "react";

// Dummy Data
const RECENT_BOARDS = [
    { id: 1, title: "Nexus", bgGradient: "linear(to-r, blue.400, purple.500)", isStarred: true },
    { id: 2, title: "CentrAlign AI", bgGradient: "linear(to-r, orange.400, red.500)", isStarred: false },
    { id: 3, title: "AICN", bgGradient: "linear(to-r, green.400, teal.500)", isStarred: false },
    { id: 4, title: "Learning", bgGradient: "linear(to-r, gray.600, gray.800)", isStarred: false },
];

const WORKSPACE_BOARDS = [
    { id: 1, title: "Nexus", bgGradient: "linear(to-r, blue.400, purple.500)", isStarred: true },
    { id: 2, title: "CentrAlign AI", bgGradient: "linear(to-r, orange.400, red.500)", isStarred: false },
    { id: 3, title: "AICN", bgGradient: "linear(to-r, green.400, teal.500)", isStarred: false },
    { id: 4, title: "Learning", bgGradient: "linear(to-r, gray.600, gray.800)", isStarred: false },
    { id: 5, title: "Project Alpha", bgGradient: "linear(to-r, pink.400, red.400)", isStarred: false },
    { id: 6, title: "Marketing Q4", bgGradient: "linear(to-r, yellow.400, orange.400)", isStarred: false },
];

export default function BoardsPage() {
    const [createBoardMenu, setCreateBoardMenu] = useState(false);
    return (
        <BoardsLayout>
            <Stack gap={8}>
                {/* Recently Viewed Section */}
                {/*<Box>
                    <Stack direction="row" align="center" mb={4} gap={2}>
                        <Icon as={FiClock} color="fg.muted" />
                        <Heading size="md" fontWeight="bold" color="fg.muted">Recently viewed</Heading>
                    </Stack>
                    <SimpleGrid columns={{ base: 1, sm: 2, md: 3, lg: 4 }} gap={4}>
                        {RECENT_BOARDS.map((board) => (
                            <BoardCard key={board.id} title={board.title} bgGradient={board.bgGradient} isStarred={board.isStarred} />
                        ))}
                    </SimpleGrid>
                </Box>*/}

                {/* Workspace Boards Section */}
                <Box>
                    <Stack direction="row" align="center" mb={4} gap={2}>
                        <Icon as={FiUsers} color="fg.muted" />
                        <Heading size="md" fontWeight="bold" color="fg.muted">Your workspaces</Heading>
                    </Stack>
                    <SimpleGrid columns={{ base: 1, sm: 2, md: 3, lg: 4 }} gap={4}>
                        {WORKSPACE_BOARDS.map((board) => (
                            <BoardCard key={board.id} title={board.title} bgGradient={board.bgGradient} isStarred={board.isStarred} />
                        ))}
                        {/* Create New Board Card Placeholder */}
                        <Popover.Root open={createBoardMenu} onOpenChange={(e) => setCreateBoardMenu(e.open)} positioning={{ placement: "right-start", offset: { mainAxis: 10, crossAxis: 0 } }}>
                            <Popover.Trigger asChild>
                                <Box
                                    h="100px"
                                    bg="gray.100"
                                    _dark={{ bg: "gray.700" }}
                                    borderRadius="md"
                                    display="flex"
                                    alignItems="center"
                                    justifyContent="center"
                                    cursor="pointer"
                                    _hover={{ bg: "gray.200", _dark: { bg: "gray.600" } }}
                                    transition="background 0.2s"
                                >
                                    <Text fontSize="sm" color="fg.muted">Create new board</Text>
                                </Box>
                            </Popover.Trigger>
                            <Portal>
                                <Popover.Positioner>
                                    <Popover.Content width="auto" p={0} borderRadius="md" boxShadow="lg">
                                        <Popover.Body p={0}>
                                            <CreateBoard onClose={() => { setCreateBoardMenu(false) }} />
                                        </Popover.Body>
                                    </Popover.Content>
                                </Popover.Positioner>
                            </Portal>
                        </Popover.Root>

                    </SimpleGrid>
                </Box>
            </Stack>
        </BoardsLayout >
    );
}
