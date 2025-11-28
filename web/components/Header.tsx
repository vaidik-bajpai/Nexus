import { Flex, Text, Button, Input, HStack, Icon, Box, Avatar, Portal, Popover } from "@chakra-ui/react";
import { FiGrid, FiBell, FiHelpCircle } from "react-icons/fi";
import { useUserStore } from "@/lib/store/auth";
import CreateBoard from "./CreateBoard";
import { useState } from "react";

export default function Header() {
    const user = useUserStore((state) => state.user);
    const [createBoardMenu, setCreateBoardMenu] = useState(false);

    return (
        <Flex
            as="header"
            align="center"
            justify="space-between"
            w="full"
            px={2}
            py={4}
            color="white"
            bg={"black"}
            position={"sticky"}
            top={0}
            zIndex={1000}
        >
            {/* Left Side */}
            <HStack gap={1}>
                <HStack gap={2} px={2}>
                    <Icon as={FiGrid} color="blue.400" />
                    <Text fontWeight="bold" fontSize="lg">Nexus</Text>
                </HStack>
            </HStack>

            <HStack>
                <Input
                    className="w-64"
                    placeholder="Search"
                    size="xs"
                    w={"3xl"}
                    bg="gray.800"
                    border="1px solid"
                    borderColor="whiteAlpha.300"
                    _hover={{ bg: "gray.700" }}
                    _focus={{ bg: "white", color: "black", borderColor: "blue.500" }}
                    color="white"
                />
                <Popover.Root open={createBoardMenu} onOpenChange={(e) => setCreateBoardMenu(e.open)}>
                    <Popover.Trigger asChild>
                        <Button
                            size="xs"
                            colorPalette="blue"
                            variant="solid"
                            px={4}
                        >
                            Create
                        </Button>
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
            </HStack>

            {/* Right Side */}
            <HStack gap={2}>
                <Button variant="ghost" size="sm" color="white" _hover={{ bg: "whiteAlpha.200" }} borderRadius="full" p={2}>
                    <Icon as={FiBell} />
                </Button>

                <Button variant="ghost" size="sm" color="white" _hover={{ bg: "whiteAlpha.200" }} borderRadius="full" p={2}>
                    <Icon as={FiHelpCircle} />
                </Button>

                <Avatar.Root size="sm" bg="orange.500" color="white" cursor="pointer">
                    <Avatar.Fallback name={user?.username || "User"} />
                </Avatar.Root>
            </HStack>
        </Flex>
    );
}
