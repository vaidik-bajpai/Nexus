import { CardMember } from "@/lib/types/cards.types";
import { Avatar, Box, Button, Flex, Heading, Icon, Input, Popover, Portal, Text } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { FiX } from "react-icons/fi";
import * as cardService from "@/lib/services/cards";

interface ChangeMembersProps {
    members: CardMember[];
    cardID: string;
    listID: string;
    boardID: string;
    onUpdate?: () => void;
}

const ChangeMembers = ({ members, cardID, listID, boardID, onUpdate }: ChangeMembersProps) => {
    const [cardMembers, setCardMembers] = useState<CardMember[] | null>(null)
    const [boardMembers, setBoardMembers] = useState<CardMember[] | null>(null)

    useEffect(() => {
        setCardMembers(members.filter((member) => member.isCardMember))
        setBoardMembers(members.filter((member) => !member.isCardMember))
    }, [members])

    async function toggleMemberToCard(member: CardMember) {
        await cardService.toggleMemberToCard({ cardID, userID: member.userID, listID, boardID })
        if (onUpdate) onUpdate()
    }

    return (
        <Portal>
            <Popover.Positioner>
                <Popover.Content width="auto" p={0} borderRadius="md" boxShadow="lg" zIndex={1600} onMouseDown={(e) => e.stopPropagation()}>
                    <Popover.Body p={0}>
                        <Box w={"xs"} color="gray.700" >
                            <Flex align="center" justify="space-between" mb={4} px={2} mt={2}>
                                <Box w={8} />
                                <Heading size="sm" fontWeight="semibold" flex={1} textAlign="center">Change members</Heading>
                                <Popover.CloseTrigger as={"div"}>
                                    <Button
                                        size="xs"
                                        variant="ghost"
                                        color="gray.500"
                                        _hover={{ color: "gray.800" }}
                                        p={0}
                                        minW={8}
                                    >
                                        <Icon as={FiX} boxSize={4} />
                                    </Button>
                                </Popover.CloseTrigger>
                            </Flex>

                            <Box mb={4} mx={3}>
                                <Input placeholder="Search members" />
                            </Box>

                            {cardMembers && cardMembers.length > 0 && <Box mx={4} mb={4}>
                                <Text fontSize="sm" fontWeight="semibold" mb={2}>Card Members</Text>
                                {cardMembers?.map((member) => {
                                    return (
                                        <Flex align="center" justify={"space-between"} key={member.userID} gap={2} _hover={{ bg: "gray.600" }} p={0.5} cursor="pointer" onClick={() => toggleMemberToCard(member)}>
                                            <Flex align="center" key={member.userID} gap={2} _hover={{ bg: "gray.600" }} p={0.5} cursor="pointer">
                                                <Avatar.Root>
                                                    <Avatar.Fallback name={member.username} />
                                                </Avatar.Root>
                                                <Text fontSize="md" color="gray.600">{member.username}</Text>
                                            </Flex>
                                            <Icon as={FiX} boxSize={4} />
                                        </Flex>
                                    )
                                })}
                            </Box>}

                            {boardMembers && boardMembers.length > 0 && <Box mx={4} mb={4}>
                                <Text fontSize="sm" fontWeight="semibold" mb={2}>Board Members</Text>
                                {boardMembers.map((member) => {
                                    return (
                                        <Flex align="center" key={member.userID} gap={2} _hover={{ bg: "gray.800" }} p={0.5} cursor="pointer" onClick={() => toggleMemberToCard(member)}>
                                            <Avatar.Root>
                                                <Avatar.Fallback name={member.username} />
                                            </Avatar.Root>
                                            <Text fontSize="md" color="gray.500">{member.username}</Text>
                                        </Flex>
                                    )
                                })}
                            </Box>}
                        </Box>
                    </Popover.Body>
                </Popover.Content>
            </Popover.Positioner>
        </Portal>
    );
};

export default ChangeMembers;