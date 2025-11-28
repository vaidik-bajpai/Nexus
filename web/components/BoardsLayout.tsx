import { Flex, Box } from "@chakra-ui/react";
import Sidebar from "./Sidebar";
import Header from "./Header";

export default function BoardsLayout({ children }: { children: React.ReactNode }) {
    return (
        <Box minH="100vh" w="full">
            <Header />
            <Flex className="w-full">
                <Sidebar />
                <div className="w-full flex justify-center overflow-hidden">
                    <Box className="max-w-7xl" w="full" flex="1" overflowY="auto" p={8}>
                        {children}
                    </Box>
                </div>
            </Flex>
        </Box>
    );
}
