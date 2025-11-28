"use client"

import { Box, Text, Flex } from "@chakra-ui/react";
import { FiStar } from "react-icons/fi";

interface BoardCardProps {
    title: string;
    bgGradient?: string;
    bgColor?: string;
    onClick: () => void;
}

export default function BoardCard({ title, bgGradient, bgColor, onClick }: BoardCardProps) {
    return (
        <Box
            h="100px"
            bgGradient={bgGradient}
            bg={bgColor}
            bgSize="cover"
            backgroundPosition="center"
            borderRadius="md"
            p={3}
            position="relative"
            cursor="pointer"
            _hover={{ filter: "brightness(0.9)" }}
            transition="filter 0.2s"
            border="1px solid rgba(255, 255, 255, 0.1)"
            onClick={onClick}
        >
            <Text fontWeight="bold" color="white" fontSize="md" textShadow="0 1px 2px rgba(0,0,0,0.2)">
                {title}
            </Text>

            <Box position="absolute" bottom={2} right={2}>
                <FiStar color="transparent" stroke="white" />
            </Box>
        </Box>
    );
}
