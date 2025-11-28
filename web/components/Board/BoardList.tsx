import { Box, Flex, Text, IconButton, HStack, Icon } from "@chakra-ui/react";
import { List } from "@/lib/types/list.types";
import BoardCard from "./BoardCard";
import { FiMoreHorizontal, FiPlus } from "react-icons/fi";

interface BoardListProps {
    list: List;
}

export default function BoardList({ list }: BoardListProps) {
    return (
        <Box
            w="272px"
            minW="272px"
            bg="blackAlpha.800"
            backdropFilter="blur(4px)"
            borderRadius="xl"
            border="1px solid"
            borderColor="whiteAlpha.100"
            p={2}
            mr={3}
            h="fit-content"
            maxH="full"
            display="flex"
            flexDirection="column"
        >
            <Flex align="center" justify="space-between" mb={2} px={2}>
                <Text fontWeight="bold" fontSize="sm" color="white">
                    {list.name}
                </Text>
                <IconButton
                    aria-label="List actions"
                    variant="ghost"
                    size="xs"
                    color="gray.400"
                    _hover={{ bg: "gray.700", color: "white" }}
                >
                    <Icon as={FiMoreHorizontal} />
                </IconButton>
            </Flex>

            <Box flex={1} overflowY="auto" px={1} className="custom-scrollbar">
                {list.cards.map((card) => (
                    <BoardCard key={card.id} card={card} />
                ))}
            </Box>

            <HStack
                mt={2}
                px={2}
                py={1.5}
                cursor="pointer"
                borderRadius="md"
                _hover={{ bg: "gray.700" }}
                color="gray.400"
                transition="background 0.2s"
            >
                <Icon as={FiPlus} />
                <Text fontSize="sm">Add a card</Text>
            </HStack>
        </Box>
    );
}
