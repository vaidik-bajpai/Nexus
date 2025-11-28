"use client"

import { Box, Heading, SimpleGrid, Text, Stack, Icon, Popover, Button, Portal } from "@chakra-ui/react";
import BoardsLayout from "@/components/BoardsLayout";
import BoardCard from "@/components/BoardCard";
import { FiClock, FiUsers } from "react-icons/fi";
import CreateBoard from "@/components/CreateBoard";
import { useState, useEffect } from "react";
import * as boardService from "@/lib/services/board"
import { Board } from "@/lib/types/board.types";
import { useRouter } from "next/navigation";


export default function BoardsPage() {
    const [createBoardMenu, setCreateBoardMenu] = useState(false);
    const [boards, setBoards] = useState<Board[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const router = useRouter();

    const fetchBoards = async () => {
        try {
            setIsLoading(true);
            const response = await boardService.listBoard({ page: 1, size: 10 });
            console.log(response);
            setBoards(response.data.boards || []);
        } catch (err) {
            console.error("Failed to fetch boards:", err);
            setError("Failed to load boards");
        } finally {
            setIsLoading(false);
        }
    };

    useEffect(() => {
        fetchBoards();
    }, []);
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
                        {isLoading ? (
                            <Text>Loading boards...</Text>
                        ) : error ? (
                            <Text color="red.500">{error}</Text>
                        ) : (
                            boards.map((board) => (
                                <BoardCard
                                    key={board.id}
                                    title={board.name}
                                    bgGradient={board.background.startsWith('#') ? undefined : `url(${board.background})`}
                                    bgColor={board.background.startsWith('#') ? board.background : undefined}
                                    onClick={() => router.push(`/boards/${board.id}`)}
                                />
                            ))
                        )}
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
                                            <CreateBoard onClose={() => {
                                                setCreateBoardMenu(false);
                                                fetchBoards();
                                            }} />
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
