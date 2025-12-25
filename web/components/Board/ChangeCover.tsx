import { Box, Button, Flex, Heading, Icon, Input, Popover, Portal, Grid, GridItem, Text, VStack } from "@chakra-ui/react";
import { FiX, FiUpload, FiImage } from "react-icons/fi";
import { useState, useRef } from "react";
import { toaster } from "@/components/ui/toaster";
import apiClient from "@/lib/apiClient";

interface ChangeCoverProps {
    onClose: () => void;
    onUpdate: (field: string, value: any) => void;
    currentCover?: string;
    currentSize?: string;
}

const COLORS = [
    "#219653", "#F2C94C", "#F2994A", "#EB5757", "#9B51E0",
    "#2F80ED", "#56CCF2", "#6FCF97", "#BB6BD9", "#828282"
];

const UNSPLASH_PHOTOS = [
    "https://images.unsplash.com/photo-1507187632231-5beb21a654a2?q=80&w=1201&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    "https://images.unsplash.com/photo-1501854140801-50d01698950b?q=80&w=1575&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    "https://plus.unsplash.com/premium_photo-1667587245819-2bea7a93e7a1?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    "https://images.unsplash.com/photo-1500673922987-e212871fec22?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
];

const ChangeCover = ({ onClose, onUpdate, currentCover, currentSize = "normal" }: ChangeCoverProps) => {
    const fileInputRef = useRef<HTMLInputElement>(null);
    const [uploading, setUploading] = useState(false);

    const handleColorSelect = (color: string) => {
        onUpdate("cover", color);
    };

    const handleSizeSelect = (size: string) => {
        onUpdate("coverSize", size);
    };

    const handlePhotoSelect = (url: string) => {
        onUpdate("cover", url);
    };

    const handleFileUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
        const file = e.target.files?.[0];
        if (!file) return;

        setUploading(true);
        const formData = new FormData();
        formData.append("file", file);

        try {
            const response = await apiClient.post("/upload", formData, {
                headers: {
                    "Content-Type": "multipart/form-data",
                },
            });
            console.log(response.data)
            onUpdate("cover", response.data.data.url);
        } catch (error) {
            console.error("Upload failed", error);
            toaster.create({
                title: "Upload failed",
                type: "error",
            });
        } finally {
            setUploading(false);
        }
    };

    return (
        <Portal>
            <Popover.Positioner>
                <Popover.Content width="auto" p={0} borderRadius="md" boxShadow="lg" zIndex={1600} onMouseDown={(e) => e.stopPropagation()}>
                    <Popover.Body p={0}>
                        <Box w="xs" color="gray.700">
                            <Flex align="center" justify="space-between" mb={2} px={2} mt={2} borderColor="gray.200" pb={2}>
                                <Box w={8} />
                                <Heading size="sm" fontWeight="semibold" flex={1} textAlign="center">Cover</Heading>
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

                            <Box px={3} pb={3} overflowY="auto">
                                <Box mb={4}>
                                    <Text fontSize="xs" fontWeight="semibold" color="gray.500" mb={2}>Size</Text>
                                    <Grid templateColumns="repeat(2, 1fr)" gap={2} h="60px">
                                        <Box
                                            borderWidth="2px"
                                            borderColor={currentSize === "normal" ? "blue.500" : "gray.200"}
                                            borderRadius="md"
                                            p={1}
                                            cursor="pointer"
                                            onClick={() => handleSizeSelect("normal")}
                                            opacity={currentSize === "normal" ? 1 : 0.6}
                                            _hover={{ opacity: 1 }}
                                        >
                                            <Box bg="gray.300" h="20px" borderRadius="sm" mb={1} />
                                            <Box bg="gray.200" h="4px" w="80%" borderRadius="sm" mb={1} />
                                            <Box bg="gray.200" h="4px" w="60%" borderRadius="sm" />
                                        </Box>
                                        <Box
                                            borderWidth="2px"
                                            borderColor={currentSize === "full" ? "blue.500" : "gray.200"}
                                            borderRadius="md"
                                            p={0}
                                            cursor="pointer"
                                            onClick={() => handleSizeSelect("full")}
                                            opacity={currentSize === "full" ? 1 : 0.6}
                                            _hover={{ opacity: 1 }}
                                            overflow="hidden"
                                            position="relative"
                                        >
                                            <Box bg="gray.300" w="full" h="full" />
                                            <Box position="absolute" bottom={1} left={1} right={1}>
                                                <Box bg="whiteAlpha.600" h="4px" w="80%" borderRadius="sm" mb={1} />
                                                <Box bg="whiteAlpha.600" h="4px" w="60%" borderRadius="sm" />
                                            </Box>
                                        </Box>
                                    </Grid>
                                </Box>

                                {currentCover && <Box>
                                    <Button
                                        w="full"
                                        size="sm"
                                        variant="outline"
                                        color="gray.600"
                                        onClick={() => onUpdate("cover", "")}
                                        mb={2}
                                    >
                                        Remove cover
                                    </Button>
                                </Box>}

                                <Box mb={4}>
                                    <Text fontSize="xs" fontWeight="semibold" color="gray.500" mb={2}>Colors</Text>
                                    <Grid templateColumns="repeat(5, 1fr)" gap={2}>
                                        {COLORS.map((color) => (
                                            <Box
                                                key={color}
                                                bg={color}
                                                h="32px"
                                                borderRadius="sm"
                                                cursor="pointer"
                                                onClick={() => handleColorSelect(color)}
                                                _hover={{ opacity: 0.8 }}
                                                borderWidth={currentCover === color ? "2px" : "0"}
                                                borderColor="blue.500"
                                            />
                                        ))}
                                    </Grid>
                                </Box>

                                <Box mb={4}>
                                    <Text fontSize="xs" fontWeight="semibold" color="gray.500" mb={2}>Attachments</Text>
                                    <Button
                                        w="full"
                                        size="sm"
                                        variant="outline"
                                        color="gray.600"
                                        onClick={() => fileInputRef.current?.click()}
                                        loading={uploading}
                                    >
                                        Upload a cover image
                                    </Button>
                                    <Input
                                        type="file"
                                        ref={fileInputRef}
                                        display="none"
                                        accept="image/*"
                                        onChange={handleFileUpload}
                                    />
                                </Box>

                                <Box>
                                    <Text fontSize="xs" fontWeight="semibold" color="gray.500" mb={2}>Photos from Unsplash</Text>
                                    <Grid templateColumns="repeat(3, 1fr)" gap={2}>
                                        {UNSPLASH_PHOTOS.map((url) => (
                                            <Button
                                                key={url}
                                                w="full"
                                                p={0}
                                                bgImage={`url(${url})`}
                                                bgSize="cover"
                                                borderRadius="sm"
                                                _hover={{ opacity: 0.8 }}
                                                onClick={() => handlePhotoSelect(url)}
                                            >
                                            </Button>
                                        ))}
                                    </Grid>
                                </Box>
                            </Box>
                        </Box>
                    </Popover.Body>
                </Popover.Content>
            </Popover.Positioner>
        </Portal>
    );
};

export default ChangeCover;
