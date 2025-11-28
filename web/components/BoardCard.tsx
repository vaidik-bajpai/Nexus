"use client"

import { Box, Text, Flex } from "@chakra-ui/react";
import { FiStar } from "react-icons/fi";

interface BoardCardProps {
    title: string;
    bgGradient: string;
    isStarred?: boolean;
}

export default function BoardCard({ title, bgGradient, isStarred }: BoardCardProps) {
    return (
        <Box
            h="100px"
            bgGradient={bgGradient}
            borderRadius="md"
            p={3}
            position="relative"
            cursor="pointer"
            _hover={{ filter: "brightness(0.9)" }}
            transition="filter 0.2s"
            border="1px solid rgba(255, 255, 255, 0.1)"
        >
            <Text fontWeight="bold" color="white" fontSize="md" textShadow="0 1px 2px rgba(0,0,0,0.2)">
                {title}
            </Text>

            <Box position="absolute" bottom={2} right={2}>
                <FiStar color={isStarred ? "#FFC107" : "transparent"} stroke={isStarred ? "#FFC107" : "white"} fill={isStarred ? "#FFC107" : "none"} />
            </Box>
        </Box>
    );
}
