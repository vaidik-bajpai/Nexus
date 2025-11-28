import { Box, Text, Image, Badge, HStack, Icon } from "@chakra-ui/react";
import { Card } from "@/lib/types/cards.types";
import { FiCheckSquare, FiAlignLeft } from "react-icons/fi";

interface BoardCardProps {
    card: Card;
}

export default function BoardCard({ card }: BoardCardProps) {
    return (
        <Box
            bg="gray.700"
            p={2}
            borderRadius="md"
            boxShadow="sm"
            mb={2}
            cursor="pointer"
            _hover={{ bg: "gray.600", boxShadow: "md", borderColor: "blue.400" }}
            border="1px solid"
            borderColor="whiteAlpha.100"
            transition="all 0.2s"
        >
            {card.cover && (
                <Image
                    src={card.cover}
                    alt={card.title}
                    borderRadius="sm"
                    mb={2}
                    objectFit="cover"
                    w="full"
                    h="120px"
                />
            )}
            <Text fontSize="sm" color="white" mb={1}>
                {card.title}
            </Text>
            <HStack gap={2} mt={1}>
                {card.completed && (
                    <Badge colorPalette="green" variant="solid" size="xs">
                        <Icon as={FiCheckSquare} mr={1} />
                        Done
                    </Badge>
                )}
                {/* Placeholder for description indicator if we had that data */}
                {/* <Icon as={FiAlignLeft} color="gray.400" size="xs" /> */}
            </HStack>
        </Box>
    );
}
