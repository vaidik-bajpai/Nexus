import { useState, useRef, useEffect } from "react";
import { Box, Button, Input, Flex, Text, IconButton, Icon } from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { FiPlus, FiX } from "react-icons/fi";
import { createList } from "@/lib/services/list";
import { toaster } from "@/components/ui/toaster";

const schema = z.object({
    name: z.string().min(1, "List name is required"),
});

type FormValues = z.infer<typeof schema>;

interface CreateListProps {
    boardId: string;
    onListCreated: () => void;
}

export default function CreateList({ boardId, onListCreated }: CreateListProps) {
    const [isEditing, setIsEditing] = useState(false);
    const formRef = useRef<HTMLDivElement>(null);
    const inputRef = useRef<HTMLInputElement>(null);

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
            await createList({ name: data.name, boardID: boardId });
            toaster.create({
                title: "List created",
                type: "success",
            });
            setIsEditing(false);
            reset();
            onListCreated();
        } catch (error) {
            toaster.create({
                title: "Failed to create list",
                type: "error",
            });
        }
    };

    if (!isEditing) {
        return (
            <Flex
                onClick={() => setIsEditing(true)}
                align="center"
                gap={2}
                w="272px"
                minW="272px"
                bg="whiteAlpha.300"
                backdropFilter="blur(4px)"
                borderRadius="xl"
                border="1px solid"
                borderColor="whiteAlpha.100"
                p={2}
                mr={3}
                color="white"
                h="fit-content"
                cursor="pointer"
                _hover={{ bg: "blackAlpha.400" }}
                transition="background 0.2s"
            >
                <Icon as={FiPlus} />
                <Text fontWeight="medium" fontSize="sm">Add another list</Text>
            </Flex>
        );
    }

    return (
        <Box
            ref={formRef}
            w="272px"
            minW="272px"
            bg="black"
            borderRadius="xl"
            p={2}
            mr={3}
            h="fit-content"
            border="1px solid"
            borderColor="blue.500"
        >
            <form onSubmit={handleSubmit(onSubmit)}>
                <Input
                    {...register("name")}
                    ref={(e) => {
                        register("name").ref(e);
                        inputRef.current = e;
                    }}
                    placeholder="Enter list name..."
                    size="sm"
                    bg="gray.800"
                    border="1px solid"
                    borderColor="whiteAlpha.300"
                    color="white"
                    mb={2}
                    _focus={{ borderColor: "blue.500", bg: "gray.900" }}
                />
                <Flex align="center" gap={2}>
                    <Button
                        type="submit"
                        size="sm"
                        colorPalette="blue"
                        loading={isSubmitting}
                    >
                        Add list
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
