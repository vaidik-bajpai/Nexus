"use client"

import { Box, Text, Flex } from "@chakra-ui/react";
import { FiStar } from "react-icons/fi";

interface BoardCardProps {
    title: string;
    bgGradient?: string;
    bgColor?: string;
}

export default function BoardCard({ title, bgGradient, bgColor }: BoardCardProps) {
    return (
        <Flex
            direction="column"
            h="100px"
            borderRadius="md"
            overflow="hidden"
            cursor="pointer"
            _hover={{ opacity: 0.9 }}
            transition="opacity 0.2s"
        >
            <Box
                flex="1"
                bgGradient={bgGradient}
                bg={bgColor}
                bgSize="cover"
                backgroundPosition="center"
            />
            <Box bg="gray.900" py={2} px={3}>
                <Text fontWeight="semibold" color="white" fontSize="sm" truncate>
                    {title}
                </Text>
            </Box>
        </Flex>
    );
}
