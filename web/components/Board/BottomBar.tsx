import { Flex, Text, HStack, Button, Icon, Separator } from "@chakra-ui/react";
import { FiStar, FiUsers, FiMoreHorizontal, FiTrello, FiActivity } from "react-icons/fi";

interface BottomBarProps {
    title: string;
}

export default function BottomBar({ title }: BottomBarProps) {
    return (
        <Flex
            as="nav"
            align="center"
            justify="space-between"
            w="full"
            px={4}
            py={2}
            bg="blackAlpha.600"
            backdropFilter="blur(10px)"
            borderBottom="1px solid"
            borderColor="whiteAlpha.200"
            color="white"
        >
            <HStack gap={4}>
                <Text fontWeight="bold" fontSize="lg">{title}</Text>
                <Button size="xs" variant="ghost" color="white" _hover={{ bg: "whiteAlpha.300" }}>
                    <Icon as={FiStar} />
                </Button>
                <Separator orientation="vertical" height="16px" borderColor="whiteAlpha.400" />
                <Button size="xs" variant="subtle" colorPalette="whiteAlpha" bg="whiteAlpha.300" _hover={{ bg: "whiteAlpha.400" }}>
                    <Icon as={FiTrello} mr={1} />
                    Board
                </Button>
            </HStack>

            <HStack gap={2}>
                <Button size="xs" variant="ghost" color="white" _hover={{ bg: "whiteAlpha.300" }}>
                    <Icon as={FiActivity} mr={1} />
                    Automation
                </Button>
                <Separator orientation="vertical" height="16px" borderColor="whiteAlpha.400" />
                <Button size="xs" variant="ghost" color="white" _hover={{ bg: "whiteAlpha.300" }}>
                    <Icon as={FiUsers} mr={1} />
                    Share
                </Button>
                <Button size="xs" variant="ghost" color="white" _hover={{ bg: "whiteAlpha.300" }}>
                    <Icon as={FiMoreHorizontal} />
                </Button>
            </HStack>
        </Flex>
    );
}
