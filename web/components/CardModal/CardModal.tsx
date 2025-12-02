import {
    DialogRoot,
    DialogContent,
    DialogHeader,
    DialogBody,
    DialogTitle,
    DialogCloseTrigger,
} from "@/components/ui/dialog"
import { Box, Flex, Text, Button, Textarea, Icon, HStack, VStack, Separator, Dialog } from "@chakra-ui/react"
import { Card, CardDetail } from "@/lib/types/cards.types"
import { useState, useEffect } from "react"
import ReactMarkdown from "react-markdown"
import remarkGfm from "remark-gfm"
import { FiList, FiCheckCircle, FiCircle } from "react-icons/fi"
import { updateCard } from "@/lib/services/cards"
import { toaster } from "@/components/ui/toaster"
import { Megaphone, Image, Ellipsis, X, Plus, Calendar, Tag, Check, User } from "lucide-react"
import * as cardsService from "@/lib/services/cards"
import CardEditButton from "./CardEditButton"
import ChangeCover from "../Board/ChangeCover"
import CardActionButton from "./CardActionButton"
import ChangeMembers from "../Board/ChangeMembers"
import AddCheckList from "../Board/AddCheckList"
import LabelPortal from "../LabelPortal"

interface CardModalProps {
    isOpen: boolean
    onClose: () => void
    card: Card
    listName: string
    boardId: string
    listId: string
    onUpdate: () => void
}

export default function CardModal({ isOpen, onClose, card, listName, boardId, listId, onUpdate }: CardModalProps) {
    const [description, setDescription] = useState(card.description || "")
    const [cardDetail, setCardDetail] = useState<CardDetail | null>(null)
    const [cover, setCover] = useState(card.cover || "")
    const [coverSize, setCoverSize] = useState(card.coverSize || "")
    const [isEditingDesc, setIsEditingDesc] = useState(false)
    const [showActivity, setShowActivity] = useState(true)

    const handleCoverUpdate = async (cover: string, coverSize: string) => {
        try {
            await updateCard({
                cardID: card.id,
                listID: listId,
                boardID: boardId,
                cover,
                coverSize
            });
            onUpdate();
        } catch (error) {
            toaster.create({
                title: "Failed to update cover",
                type: "error",
            });
        }
    };

    useEffect(() => {
        cardsService.getCardDetail({
            cardId: card.id,
            boardId: boardId,
            listId: listId
        }).then((res) => {
            setCardDetail(res.data)
        })
    }, [card])

    const handleSaveDescription = async () => {
        try {
            await cardsService.updateCard({
                cardID: card.id,
                listID: listId,
                boardID: boardId,
                description
            })
            onUpdate()
            setIsEditingDesc(false)
        } catch (error) {
            toaster.create({
                title: "Failed to update description",
                type: "error"
            })
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
                            <CardEditButton icon={<Image size={16} />} tooltipText="Add image" portal={<ChangeCover onClose={() => { }} onUpdate={handleCoverUpdate} currentCover={cover} currentSize={coverSize} />} />
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
                                _hover={{ color: "green.400" }}
                            />
                        ) : (
                            <Icon
                                as={FiCheckCircle}
                                color="green.400"
                                boxSize={4}
                                mt={0.5}
                                animation="fadeIn 0.3s ease-out 0.1s both"
                                cursor="pointer"
                            />
                        )}
                        <Box>
                            <DialogTitle fontSize="2xl" fontWeight="bold">{card.title}</DialogTitle>
                        </Box>
                    </Flex>
                </DialogHeader>

                <DialogBody>
                    <Flex gap={2}>
                        <CardActionButton icon={<Plus />} text="Add" />
                        <CardActionButton icon={<Tag />} text="Label" portal={<LabelPortal boardId={boardId} />} />
                        <CardActionButton icon={<Calendar />} text="Dates" />
                        <CardActionButton icon={<Check />} text="Checklist" portal={<AddCheckList />} />
                        <CardActionButton icon={<User />} text="Members" portal={<ChangeMembers members={cardDetail?.members || []} cardID={card.id} listID={listId} boardID={boardId} onUpdate={onUpdate} />} />
                    </Flex>
                </DialogBody>

                <DialogBody>
                    <Flex gap={8} direction={{ base: "column", md: "row" }}>
                        <Box flex={1}>
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
                                            <Button size="sm" colorPalette="blue" onClick={handleSaveDescription}>Save</Button>
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

                            {/* Activity Section */}
                            {/* <Flex align="center" justify="space-between" mb={4}>
                                <Flex align="center" gap={4}>
                                    <Icon as={FiActivity} boxSize={6} color="gray.400" />
                                    <Text fontSize="lg" fontWeight="semibold">Activity</Text>
                                </Flex>
                                <Button size="sm" variant="subtle" onClick={() => setShowActivity(!showActivity)}>
                                    {showActivity ? "Hide details" : "Show details"}
                                </Button>
                            </Flex>
                            */}
                            {
                                // showActivity && (
                                //     <Box ml={10}>
                                //         <Flex gap={2} mb={4}>
                                //             <Box w={8} h={8} borderRadius="full" bg="orange.500" display="flex" alignItems="center" justifyContent="center" fontWeight="bold">
                                //                 VB
                                //             </Box>
                                //             <Box flex={1}>
                                //                 <Textarea
                                //                     placeholder="Write a comment..."
                                //                     bg="gray.800"
                                //                     border="none"
                                //                     rows={1}
                                //                     resize="none"
                                //                     _focus={{ ring: 2, ringColor: "blue.500", minH: "80px" }}
                                //                     transition="all 0.2s"
                                //                 />
                                //             </Box>
                                //         </Flex>
                                //         {/* Placeholder for activity log */}
                                //         <VStack align="stretch" gap={4}>
                                //             <Flex gap={3}>
                                //                 <Box w={8} h={8} borderRadius="full" bg="orange.500" display="flex" alignItems="center" justifyContent="center" fontSize="xs" fontWeight="bold">
                                //                     VB
                                //                 </Box>
                                //                 <Box>
                                //                     <Text fontSize="sm"><Text as="span" fontWeight="bold">Vaidik Bajpai</Text> added this card to {listName}</Text>
                                //                     <Text fontSize="xs" color="gray.500">Just now</Text>
                                //                 </Box>
                                //             </Flex>
                                //         </VStack>
                                //     </Box>
                                // )
                            }
                        </Box>
                    </Flex>
                </DialogBody>
            </DialogContent>
        </DialogRoot >
    )
}
