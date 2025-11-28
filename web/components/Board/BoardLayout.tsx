import { Box, Flex } from "@chakra-ui/react";
import Header from "../Header";
import BottomBar from "./BottomBar";
import FloatingBottomNav from "./FloatingBottomNav";
import { ReactNode } from "react";

interface BoardLayoutProps {
    children: ReactNode;
    background: string;
    title: string;
}

export default function BoardLayout({ children, background, title }: BoardLayoutProps) {
    const isImage = background.startsWith("http") || background.startsWith("/") || background.startsWith("url");
    const bgImage = isImage ? (background.startsWith("url") ? background : `url('${background}')`) : undefined;
    const bgColor = isImage ? "gray.900" : background;

    return (
        <Box
            h="100vh"
            w="100vw"
            bg={bgColor}
            bgImage={bgImage}
            bgSize="cover"
            backgroundPosition="center"
            bgRepeat="no-repeat"
            display="flex"
            flexDirection="column"
            overflow="hidden"
            position="relative"
        >
            <Header />
            <Flex flex={1} direction="column" overflow="hidden">
                <BottomBar title={title} />
                <Box flex={1} overflowX="auto" overflowY="hidden" p={4}>
                    {children}
                </Box>
            </Flex>
            <FloatingBottomNav />
        </Box>
    );
}
