"use client"

import { Box, Button, Flex, Grid, Heading, Input, Popover, Text, VStack, Icon, Image, Center } from "@chakra-ui/react";
import { useState } from "react";
import { FiX, FiCheck, FiMoreHorizontal } from "react-icons/fi";

const BACKGROUNDS = [
    { type: "image", value: "https://images.unsplash.com/photo-1707343843437-caacff5cfa74?q=80&w=100&auto=format&fit=crop", id: "bg1" },
    { type: "image", value: "https://images.unsplash.com/photo-1682687220742-aba13b6e50ba?q=80&w=100&auto=format&fit=crop", id: "bg2" },
    { type: "image", value: "https://images.unsplash.com/photo-1707345512638-997d31a10eaa?q=80&w=100&auto=format&fit=crop", id: "bg3" },
    { type: "image", value: "https://images.unsplash.com/photo-1707343848552-893e05dba6ac?q=80&w=100&auto=format&fit=crop", id: "bg4" },
    { type: "color", value: "#0079bf", id: "blue" },
    { type: "color", value: "#d29034", id: "orange" },
    { type: "color", value: "#519839", id: "green" },
    { type: "color", value: "#b04632", id: "red" },
    { type: "color", value: "#89609e", id: "purple" },
];

export default function CreateBoard({ onClose }: { onClose: () => void }) {
    const [title, setTitle] = useState("");
    const [selectedBg, setSelectedBg] = useState(BACKGROUNDS[0].id);

    const selectedBackground = BACKGROUNDS.find(bg => bg.id === selectedBg);

    return (
        <Box w={"1xs"} color="gray.700">
            {/* Header */}
            <Flex align="center" justify="space-between" mb={4} px={1}>
                <Box w={8} /> {/* Spacer for centering */}
                <Heading size="sm" fontWeight="semibold" flex={1} textAlign="center">Create board</Heading>
                <Button
                    size="xs"
                    variant="ghost"
                    onClick={onClose}
                    color="gray.500"
                    _hover={{ color: "gray.800" }}
                    p={0}
                    minW={8}
                >
                    <Icon as={FiX} boxSize={4} />
                </Button>
            </Flex>

            {/* Preview */}
            <Box
                mx={"auto"}
                mb={4}
                w={"3xs"}
                borderRadius="md"
                h="120px"
                bg={selectedBackground?.type === 'color' ? selectedBackground.value : `url(${selectedBackground?.value})`}
                backgroundSize="cover"
                backgroundPosition="center"
                display="flex"
                alignItems="center"
                justifyContent="center"
                boxShadow="sm"
            >
                <Image src="https://trello.com/assets/14cda5dc635d1f13bc48.svg" alt="Preview" w="80%" />
            </Box>

            {/* Background Picker */}
            <Box mb={4} mx={3}>
                <Text fontSize="xs" fontWeight="bold" mb={2} color="gray.600">Background</Text>
                <Grid templateColumns="repeat(4, 1fr)" gap={2}>
                    {BACKGROUNDS.slice(0, 4).map((bg) => (
                        <Button
                            key={bg.id}
                            w="full"
                            h="40px"
                            p={0}
                            bg={bg.type === 'color' ? bg.value : `url(${bg.value})`}
                            backgroundSize="cover"
                            backgroundPosition="center"
                            borderRadius="md"
                            _hover={{ opacity: 0.8 }}
                            onClick={() => setSelectedBg(bg.id)}
                            position="relative"
                        >
                            {selectedBg === bg.id && (
                                <Center w="full" h="full" bg="blackAlpha.300">
                                    <Icon as={FiCheck} color="white" />
                                </Center>
                            )}
                        </Button>
                    ))}
                </Grid>
                <Grid templateColumns="repeat(6, 1fr)" gap={2} mt={2}>
                    {BACKGROUNDS.slice(4).map((bg) => (
                        <Button
                            key={bg.id}
                            w="full"
                            h="32px"
                            p={0}
                            bg={bg.value}
                            borderRadius="sm"
                            _hover={{ opacity: 0.8 }}
                            onClick={() => setSelectedBg(bg.id)}
                        >
                            {selectedBg === bg.id && (
                                <Center w="full" h="full" bg="blackAlpha.300">
                                    <Icon as={FiCheck} color="white" boxSize={3} />
                                </Center>
                            )}
                        </Button>
                    ))}
                    <Button w="full" h="32px" p={0} bg="gray.100" borderRadius="sm" _hover={{ bg: "gray.200" }}>
                        <Icon as={FiMoreHorizontal} color="gray.600" />
                    </Button>
                </Grid>
            </Box>

            {/* Form */}
            <VStack align="stretch" gap={4} mx={3} mb={3}>
                <Box>
                    <Text fontSize="xs" fontWeight="bold" mb={1} color="gray.600">Board title <Text as="span" color="red.500">*</Text></Text>
                    <Input
                        size="sm"
                        value={title}
                        onChange={(e) => setTitle(e.target.value)}
                        borderColor={!title ? "red.300" : "gray.300"}
                        _focus={{ borderColor: "blue.500", boxShadow: "0 0 0 1px var(--chakra-colors-blue-500)" }}
                    />
                    {!title && (
                        <Text fontSize="xs" color="gray.600" mt={1}>👋 Board title is required</Text>
                    )}
                </Box>

                <Box>
                    <Text fontSize="xs" fontWeight="bold" mb={1} color="gray.600">Visibility</Text>
                    <Button
                        size="sm"
                        variant="outline"
                        w="full"
                        justifyContent="space-between"
                        fontWeight="normal"
                        color="gray.700"
                    >
                        Workspace
                    </Button>
                </Box>

                <Button
                    colorPalette="blue"
                    width="full"
                    disabled={!title}
                >
                    Create
                </Button>

                {/* <Button variant="ghost" size="sm" color="gray.600">
                    Start with a template
                </Button> */}
            </VStack>
        </Box>
    );
}
