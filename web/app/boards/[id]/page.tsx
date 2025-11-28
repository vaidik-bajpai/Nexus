"use client"

import { useEffect, useState, use } from "react";
import { getBoard } from "@/lib/services/board";
import { BoardDetail } from "@/lib/types/board.types";
import BoardLayout from "@/components/Board/BoardLayout";
import BoardList from "@/components/Board/BoardList";
import CreateList from "@/components/Board/CreateList";
import { Flex, Spinner, Center, Text, Box, Icon } from "@chakra-ui/react";
import { FiPlus } from "react-icons/fi";

interface PageProps {
    params: Promise<{ id: string }>;
}

export default function BoardPage({ params }: PageProps) {
    const { id } = use(params);
    const [board, setBoard] = useState<BoardDetail | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
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
            <Flex h="full" align="flex-start">
                {board.lists && board.lists.map((list) => (
                    <BoardList key={list.id} list={list} boardId={board.id} onCardCreated={fetchBoard} />
                ))}
                <CreateList boardId={board.id} onListCreated={fetchBoard} />
            </Flex>
        </BoardLayout>
    );
}
