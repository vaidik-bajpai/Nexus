import {
    DialogRoot,
    DialogContent,
    DialogHeader,
    DialogBody,
    DialogTitle,
    DialogCloseTrigger,
} from "@/components/ui/dialog"
import { Box, Flex, Text, Button, Textarea, Icon, HStack, VStack, Separator, Dialog, Input, Popover } from "@chakra-ui/react"
import { Card, CardDetail } from "@/lib/types/cards.types"
import { useState, useEffect } from "react"
import ReactMarkdown from "react-markdown"
import remarkGfm from "remark-gfm"
import { FiList, FiCheckCircle, FiCircle } from "react-icons/fi"
import { updateCard } from "@/lib/services/cards"
import { toaster } from "@/components/ui/toaster"
import { Tooltip } from "@/components/ui/tooltip"
import { Megaphone, Image, Ellipsis, X, Plus, Calendar, Tag, Check, User } from "lucide-react"
import * as cardsService from "@/lib/services/cards"
import CardEditButton from "./CardEditButton"
import ChangeCover from "../Board/ChangeCover"
import CardActionButton from "./CardActionButton"
import ChangeMembers from "../Board/ChangeMembers"
import AddCheckList from "../Board/AddCheckList"
import Checklist, { ChecklistData } from "./Checklist"
import LabelPortal from "../LabelPortal"
import { useBoardStore } from "@/lib/store/board"

interface CardModalProps {
    isOpen: boolean
    onClose: () => void
    cardId: string
    listName: string
    boardId: string
    listId: string
}

export default function CardModal({ isOpen, onClose, cardId, listName, boardId, listId }: CardModalProps) {
    const { cards, enrichCards, metadata } = useBoardStore()
    const card = cards.find(c => c.id === cardId)

    const [description, setDescription] = useState(card?.description || "")
    const [isEditingDesc, setIsEditingDesc] = useState(false)
    const [title, setTitle] = useState(card?.title || "")
    const [checklists, setChecklists] = useState<ChecklistData[]>([])

    useEffect(() => {
        if (card?.title) setTitle(card.title)
        if (card?.checklist) {
            setChecklists(card.checklist.map(cl => ({
                id: crypto.randomUUID(),
                title: cl.title,
                items: cl.items.map(item => ({
                    id: item.id,
                    title: item.title,
                    done: item.done
                }))
            })));
        } else {
            setChecklists([]);
        }
    }, [card?.title, card?.checklist])

    const handleTitleSave = () => {
        if (card && title !== card.title) {
            handleUpdateCardField("title", title)
        }
    }

    const handleAddChecklist = (title: string) => {
        const newChecklist: ChecklistData = {
            id: crypto.randomUUID(),
            title,
            items: []
        }
        setChecklists([...checklists, newChecklist])
    }

    const handleDeleteChecklist = (id: string) => {
        setChecklists(checklists.filter(c => c.id !== id))
    }

    const handleUpdateChecklist = (id: string, data: Partial<ChecklistData>) => {
        setChecklists(checklists.map(c => c.id === id ? { ...c, ...data } : c))
    }

    useEffect(() => {
        if (isOpen && cardId) {
            cardsService.getCardDetail({
                cardId: cardId,
                boardId: boardId,
                listId: listId
            }).then((res) => {
                const detail = res.data;
                enrichCards(cardId, {
                    description: detail.description,
                    member_ids: detail.member_ids,
                    checklist: detail.checklist,
                    archived: detail.archived,
                    start: detail.start ? detail.start.toString() : undefined,
                    // due: detail.due ? detail.due.toString() : undefined, // Already in Card
                    labels: detail.labels,
                    completed: detail.completed,
                    cover: detail.cover,
                    coverSize: detail.coverSize
                })
            })
        }
    }, [isOpen, cardId, boardId, listId, enrichCards])

    if (!card) return null

    // Derived state
    const displayMembers = (card.member_ids || []).map(id => {
        const member = metadata?.members.find(m => m.id === id);
        return member ? {
            userID: member.id,
            username: member.username,
            fullName: member.fullName,
            email: member.email,
            avatar: "",
            isCardMember: true
        } : null;
    }).filter(Boolean) as any[];

    const handleUpdateCardField = async (fieldName: string, fieldValue: any) => {
        const oldValue = card[fieldName as keyof Card];
        try {
            enrichCards(card.id, { [fieldName]: fieldValue });
            await updateCard({
                cardID: card.id,
                listID: listId,
                boardID: boardId,
                [fieldName]: fieldValue
            });
        } catch (error) {
            enrichCards(card.id, { [fieldName]: oldValue });
            toaster.create({
                title: "Failed to update card, reverting",
                type: "error",
            });
        }
    }

    return (
        <DialogRoot open={isOpen} onOpenChange={(e) => !e.open && onClose()} size="md" placement="center" >
            <DialogContent bg="gray.900" color="white" maxW="2xl" borderRadius="xl">
                <Flex justify="space-between" align={"center"} p={4} borderBottom={"1px solid"} borderColor="gray.700">
                    <Button size="xs" variant={"surface"} w={"fit-content"} px={2} bg={"gray.700"}>
                        {listName}
                    </Button>
                    <Flex gap={6} color="gray.400">
                        <Flex gap={4}>
                            <CardEditButton icon={<Megaphone size={16} />} tooltipText="Share feedback" />
                            <Separator orientation="vertical" color="gray.400" size={"sm"} />
                            <CardEditButton icon={<Image size={16} />} tooltipText="Add image" portal={<ChangeCover onClose={() => { }} onUpdate={handleUpdateCardField} currentCover={card.cover} currentSize={card.coverSize} />} />
                        </Flex>
                        <CardEditButton icon={<Ellipsis size={16} />} tooltipText="More options" />
                        <CardEditButton icon={<X size={16} />} tooltipText="Close" onClick={onClose} />
                    </Flex>
                </Flex>
                <DialogHeader>
                    <Flex align="center" gap={4} justify="center">
                        {!card.completed ? (
                            <Icon
                                as={FiCircle}
                                color="gray.400"
                                boxSize={4}
                                mt={0.5}
                                animation="fadeIn 0.3s ease-out 0.1s both"
                                cursor="pointer"
                                onClick={() => handleUpdateCardField("completed", true)}
                                _hover={{ color: "green.400" }}
                            />
                        ) : (
                            <Icon
                                as={FiCheckCircle}
                                color="green.400"
                                boxSize={4}
                                mt={0.5}
                                animation="fadeIn 0.3s ease-out 0.1s both"
                                onClick={() => handleUpdateCardField("completed", false)}
                                cursor="pointer"
                            />
                        )}
                        <Box flex={1}>
                            <Input
                                value={title}
                                onChange={(e) => setTitle(e.target.value)}
                                onBlur={handleTitleSave}
                                onKeyDown={(e) => {
                                    if (e.key === "Enter") {
                                        e.currentTarget.blur()
                                    }
                                }}
                                fontSize="2xl"
                                fontWeight="bold"
                                px={2}
                                py={1}
                                borderRadius="md"
                                _focus={{
                                    border: "2px solid",
                                    borderColor: "blue.500",
                                    bg: "gray.800"
                                }}
                                border="2px solid"
                                borderColor="transparent"
                                bg="transparent"
                            />
                        </Box>
                    </Flex>
                </DialogHeader>

                <DialogBody>
                    <Flex gap={2}>
                        <CardActionButton icon={<Plus />} text="Add" />
                        <CardActionButton isHidden={!card.labels || card.labels.length === 0} icon={<Tag />} text="Label" portal={<LabelPortal boardId={boardId} cardId={cardId} listId={listId} activeLabels={card.labels} />} />
                        <CardActionButton icon={<Calendar />} text="Dates" />
                        <CardActionButton icon={<Check />} text="Checklist" portal={<AddCheckList onAdd={handleAddChecklist} />} />
                        <CardActionButton isHidden={displayMembers.length > 0} icon={<User />} text="Members" portal={<ChangeMembers memberIds={card.member_ids || []} cardID={card.id} listID={listId} boardID={boardId} />} />
                    </Flex>
                </DialogBody>

                <DialogBody>
                    <Flex gap={8} direction={{ base: "column", md: "row" }}>
                        <Box flex={1}>
                            {/* Members and Labels Section */}
                            <Flex gap={4} mb={6} wrap="wrap">
                                {displayMembers.length > 0 && (
                                    <Box>
                                        <Text fontSize="xs" fontWeight="semibold" color="gray.400" mb={2}>Members</Text>
                                        <Flex gap={2} wrap="wrap">
                                            {displayMembers.map((member) => (
                                                <Tooltip
                                                    key={member.userID}
                                                    content={member.fullName || member.username}
                                                    positioning={{ placement: "top" }}
                                                >
                                                    <Box
                                                        bg="gray.700"
                                                        borderRadius="full"
                                                        w={8}
                                                        h={8}
                                                        display="flex"
                                                        alignItems="center"
                                                        justifyContent="center"
                                                        cursor="default"
                                                    >
                                                        <Text fontSize="xs" fontWeight="bold">{member.username.charAt(0).toUpperCase()}</Text>
                                                    </Box>
                                                </Tooltip>
                                            ))}
                                            <Popover.Root positioning={{ placement: "bottom-start" }}>
                                                <Popover.Trigger asChild>
                                                    <Box
                                                        as="button"
                                                        bg="gray.700"
                                                        w={8}
                                                        h={8}
                                                        borderRadius="full"
                                                        cursor="pointer"
                                                        _hover={{ bg: "gray.600" }}
                                                        display="flex"
                                                        alignItems="center"
                                                        justifyContent="center"
                                                    >
                                                        <Icon as={Plus} boxSize={4} color="gray.400" />
                                                    </Box>
                                                </Popover.Trigger>
                                                <ChangeMembers memberIds={card.member_ids || []} cardID={card.id} listID={listId} boardID={boardId} />
                                            </Popover.Root>
                                        </Flex>
                                    </Box>
                                )}

                                {card.labels && card.labels.length > 0 && (
                                    <Box>
                                        <Text fontSize="xs" fontWeight="semibold" color="gray.400" mb={2}>Labels</Text>
                                        <Flex gap={2} wrap="wrap">
                                            {card.labels.map((label: any) => (
                                                <Box
                                                    key={label.labelID || label.id}
                                                    bg={label.color}
                                                    px={3}
                                                    h={8}
                                                    borderRadius="sm"
                                                    cursor="pointer"
                                                    _hover={{ opacity: 0.8 }}
                                                    display="flex"
                                                    alignItems="center"
                                                >
                                                    <Text fontSize="xs" fontWeight="bold" color="white">{label.name}</Text>
                                                </Box>
                                            ))}
                                            <Popover.Root positioning={{ placement: "bottom-start" }} lazyMount unmountOnExit>
                                                <Popover.Trigger asChild>
                                                    <Box
                                                        as="button"
                                                        bg="gray.700"
                                                        w={8}
                                                        h={8}
                                                        borderRadius="sm"
                                                        cursor="pointer"
                                                        _hover={{ bg: "gray.600" }}
                                                        display="flex"
                                                        alignItems="center"
                                                        justifyContent="center"
                                                    >
                                                        <Icon as={Plus} boxSize={4} color="gray.400" />
                                                    </Box>
                                                </Popover.Trigger>
                                                <LabelPortal boardId={boardId} cardId={cardId} listId={listId} activeLabels={card.labels} />
                                            </Popover.Root>
                                        </Flex>
                                    </Box>
                                )}
                            </Flex>

                            {/* Description Section */}
                            <Flex align="center" gap={4} mb={4}>
                                <Icon as={FiList} boxSize={6} color="gray.400" />
                                <Text fontSize="lg" fontWeight="semibold">Description</Text>
                            </Flex>

                            <Box ml={10} mb={8}>
                                {isEditingDesc ? (
                                    <Box>
                                        <Textarea
                                            value={description}
                                            onChange={(e) => setDescription(e.target.value)}
                                            placeholder="Add a more detailed description..."
                                            minH="150px"
                                            bg="gray.800"
                                            border="none"
                                            _focus={{ ring: 2, ringColor: "blue.500" }}
                                            mb={2}
                                        />
                                        <HStack>
                                            <Button size="sm" colorPalette="blue" onClick={() => handleUpdateCardField("description", description)}>Save</Button>
                                            <Button size="sm" variant="ghost" onClick={() => setIsEditingDesc(false)}>Cancel</Button>
                                        </HStack>
                                    </Box>
                                ) : (
                                    <Box
                                        onClick={() => setIsEditingDesc(true)}
                                        cursor="pointer"
                                        minH="60px"
                                        bg={description ? "transparent" : "gray.800"}
                                        p={description ? 0 : 4}
                                        borderRadius="md"
                                        _hover={{ bg: description ? "gray.800" : "gray.700" }}
                                        transition="background 0.2s"
                                    >
                                        {description ? (
                                            <Box className="markdown-body" p={2}>
                                                <ReactMarkdown remarkPlugins={[remarkGfm]}>{description}</ReactMarkdown>
                                            </Box>
                                        ) : (
                                            <Text color="gray.400">Add a more detailed description...</Text>
                                        )}
                                    </Box>
                                )}
                            </Box>

                            {/* Checklists Section */}
                            {checklists.map(checklist => (
                                <Checklist
                                    key={checklist.id}
                                    checklist={checklist}
                                    onDelete={handleDeleteChecklist}
                                    onUpdate={handleUpdateChecklist}
                                />
                            ))}
                        </Box>
                    </Flex>
                </DialogBody>
            </DialogContent>
        </DialogRoot >
    )
}
