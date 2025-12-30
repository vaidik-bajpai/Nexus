import { CardMember } from "@/lib/types/cards.types";
import { Avatar, Box, Button, Flex, Heading, Icon, Input, Popover, Portal, Text } from "@chakra-ui/react";
import { useState } from "react";
import { FiX, FiCheck } from "react-icons/fi";
import * as cardService from "@/lib/services/cards";
import { useBoardStore } from "@/lib/store/board";
import { BoardMember } from "@/lib/types/board.types";

interface ChangeMembersProps {
    memberIds: string[];
    cardID: string;
    listID: string;
    boardID: string;
}

const ChangeMembers = ({ memberIds, cardID, listID, boardID }: ChangeMembersProps) => {
    const { metadata, enrichCards } = useBoardStore();
    const [searchQuery, setSearchQuery] = useState("");


    const boardMembers = metadata?.members || [];

    const assignedMembers = boardMembers.filter(bm =>
        memberIds.includes(bm.id)
    );

    const unassignedMembers = boardMembers.filter(bm =>
        !memberIds.includes(bm.id)
    ).filter(bm =>
        bm.username.toLowerCase().includes(searchQuery.toLowerCase()) ||
        bm.fullName.toLowerCase().includes(searchQuery.toLowerCase())
    );

    async function toggleMemberToCard(memberId: string) {
        const isAssigned = memberIds.includes(memberId);
        let newMemberIds = [...memberIds];

        if (isAssigned) {
            newMemberIds = newMemberIds.filter(id => id !== memberId);
        } else {
            newMemberIds.push(memberId);
        }


        enrichCards(cardID, { member_ids: newMemberIds });

        try {
            await cardService.toggleMemberToCard({ cardID, userID: memberId, listID, boardID });
        } catch (error) {
            enrichCards(cardID, { member_ids: memberIds });
        }
    }

    return (
        <Portal>
            <Box
                position="fixed"
                top="4rem"
                left="50%"
                transform="translateX(-50%)"
                zIndex={1600}
                maxH="calc(100vh - 5rem)"
                w="300px"
                outline="none"
            >
                <Popover.Content
                    width="full"
                    maxH="full"
                    overflowY="auto"
                    p={0}
                    borderRadius="md"
                    boxShadow="lg"
                    onMouseDown={(e) => e.stopPropagation()}
                    onClick={(e) => e.stopPropagation()}
                    bg="gray.800"
                    border="1px solid"
                    borderColor="gray.700"
                >
                    <Popover.Body p={0}>
                        <Box color="gray.300" px={4} pb={4} pt={2}>
                            <Flex align="center" justify="space-between" mb={4} px={0} borderBottom="1px solid" borderColor="gray.700" pb={2}>
                                <Box w={8} />
                                <Heading size="sm" fontWeight="semibold" flex={1} textAlign="center" color="gray.300">
                                    Members
                                </Heading>
                                <Popover.CloseTrigger asChild>
                                    <Button
                                        size="xs"
                                        variant="ghost"
                                        color="gray.400"
                                        _hover={{ color: "white", bg: "gray.700" }}
                                        p={0}
                                        minW={8}
                                        onClick={(e) => e.stopPropagation()}
                                    >
                                        <Icon as={FiX} boxSize={4} />
                                    </Button>
                                </Popover.CloseTrigger>
                            </Flex>

                            <Box mb={4}>
                                <Input
                                    placeholder="Search members..."
                                    value={searchQuery}
                                    onChange={(e) => setSearchQuery(e.target.value)}
                                    bg="gray.900"
                                    border="1px solid"
                                    borderColor="gray.700"
                                    _focus={{ borderColor: "blue.500", outline: "none" }}
                                    size="sm"
                                    color="white"
                                />
                            </Box>

                            <Flex direction="column" gap={1}>
                                <Text fontSize="xs" fontWeight="bold" color="gray.500" mb={1}>Card Members</Text>
                                {assignedMembers.length === 0 && <Text fontSize="xs" color="gray.500" fontStyle="italic">No members assigned</Text>}
                                {assignedMembers.map((member) => (
                                    <MemberItem
                                        key={member.id}
                                        member={member}
                                        isAssigned={true}
                                        onToggle={() => toggleMemberToCard(member.id)}
                                    />
                                ))}
                            </Flex>

                            <Flex direction="column" gap={1} mt={4}>
                                <Text fontSize="xs" fontWeight="bold" color="gray.500" mb={1}>Board Members</Text>
                                {unassignedMembers.length === 0 && <Text fontSize="xs" color="gray.500" fontStyle="italic">No other members found</Text>}
                                {unassignedMembers.map((member) => (
                                    <MemberItem
                                        key={member.id}
                                        member={member}
                                        isAssigned={false}
                                        onToggle={() => toggleMemberToCard(member.id)}
                                    />
                                ))}
                            </Flex>
                        </Box>
                    </Popover.Body>
                </Popover.Content>
            </Box>
        </Portal>
    );
};

const MemberItem = ({ member, isAssigned, onToggle }: { member: BoardMember, isAssigned: boolean, onToggle: () => void }) => {
    return (
        <Flex
            align="center"
            justify="space-between"
            gap={2}
            p={2}
            cursor="pointer"
            borderRadius="sm"
            _hover={{ bg: "gray.700" }}
            onClick={onToggle}
        >
            <Flex align="center" gap={2}>
                <Avatar.Root size="xs">
                    <Avatar.Fallback name={member.fullName} />
                </Avatar.Root>
                <Text fontSize="sm" color="gray.300">{member.fullName}</Text>
            </Flex>
            {isAssigned && <Icon as={FiCheck} boxSize={4} color="blue.400" />}
        </Flex>
    )
}

export default ChangeMembers;