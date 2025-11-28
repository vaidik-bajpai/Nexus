import { useState, useRef, useEffect } from "react";
import { Box, Button, Textarea, Flex, Text, IconButton, Icon, HStack } from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { FiPlus, FiX } from "react-icons/fi";
import { createCard } from "@/lib/services/cards";
import { toaster } from "@/components/ui/toaster";

const schema = z.object({
    title: z.string().min(1, "Card title is required"),
});

type FormValues = z.infer<typeof schema>;

interface CreateCardProps {
    listId: string;
    boardId: string;
    onCardCreated: () => void;
}

export default function CreateCard({ listId, boardId, onCardCreated }: CreateCardProps) {
    const [isEditing, setIsEditing] = useState(false);
    const formRef = useRef<HTMLDivElement>(null);
    const inputRef = useRef<HTMLTextAreaElement>(null);

    const {
        register,
        handleSubmit,
        reset,
        formState: { isSubmitting },
    } = useForm<FormValues>({
        resolver: zodResolver(schema),
    });

    useEffect(() => {
        if (isEditing && inputRef.current) {
            inputRef.current.focus();
        }
    }, [isEditing]);

    useEffect(() => {
        const handleClickOutside = (event: MouseEvent) => {
            if (formRef.current && !formRef.current.contains(event.target as Node)) {
                setIsEditing(false);
                reset();
            }
        };

        if (isEditing) {
            document.addEventListener("mousedown", handleClickOutside);
        }

        return () => {
            document.removeEventListener("mousedown", handleClickOutside);
        };
    }, [isEditing, reset]);

    const onSubmit = async (data: FormValues) => {
        try {
            await createCard({ title: data.title, listID: listId, boardID: boardId });
            toaster.create({
                title: "Card created",
                type: "success",
            });
            setIsEditing(false);
            reset();
            onCardCreated();
        } catch (error) {
            toaster.create({
                title: "Failed to create card",
                type: "error",
            });
        }
    };

    if (!isEditing) {
        return (
            <HStack
                onClick={() => setIsEditing(true)}
                mt={2}
                px={2}
                py={1.5}
                cursor="pointer"
                borderRadius="md"
                _hover={{ bg: "whiteAlpha.200" }}
                color="gray.400"
                transition="background 0.2s"
            >
                <Icon as={FiPlus} />
                <Text fontSize="sm">Add a card</Text>
            </HStack>
        );
    }

    return (
        <Box ref={formRef} mt={2} px={1}>
            <form onSubmit={handleSubmit(onSubmit)}>
                <Textarea
                    {...register("title")}
                    ref={(e) => {
                        register("title").ref(e);
                        inputRef.current = e;
                    }}
                    placeholder="Enter a title or paste a link"
                    size="sm"
                    bg="gray.800"
                    border="1px solid"
                    borderColor="blue.500"
                    color="white"
                    mb={2}
                    resize="none"
                    rows={2}
                    _focus={{ borderColor: "blue.500", bg: "gray.900" }}
                />
                <Flex align="center" gap={2}>
                    <Button
                        type="submit"
                        size="sm"
                        colorPalette="blue"
                        loading={isSubmitting}
                    >
                        Add card
                    </Button>
                    <IconButton
                        aria-label="Cancel"
                        size="sm"
                        variant="ghost"
                        color="gray.400"
                        _hover={{ color: "white", bg: "whiteAlpha.200" }}
                        onClick={() => {
                            setIsEditing(false);
                            reset();
                        }}
                    >
                        <Icon as={FiX} />
                    </IconButton>
                </Flex>
            </form>
        </Box>
    );
}
