"use client"

import { Box, Text, Flex, Icon } from "@chakra-ui/react";
import { FiAlignLeft, FiCheckSquare } from "react-icons/fi";

interface BoardCardItemProps {
    title: string;
}

export default function BoardCardItem({ title }: BoardCardItemProps) {
    return (
        <Box
            bg="gray.800"
            p={3}
            borderRadius="md"
            boxShadow="sm"
            cursor="pointer"
            _hover={{ ring: "2px", ringColor: "blue.500" }}
            mb={2}
            color="white"
        >
            <Text fontSize="sm" mb={2}>{title}</Text>

        </Box>
    );
}
