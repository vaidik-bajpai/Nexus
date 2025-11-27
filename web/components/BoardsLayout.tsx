"use client"

import { Flex, Box } from "@chakra-ui/react";
import Sidebar from "./Sidebar";

export default function BoardsLayout({ children }: { children: React.ReactNode }) {
    return (
        <Flex className="min-h-screen w-full"> {/* Adjust mt based on your Navbar height */}
            <Sidebar />
            <div className="w-full flex justify-center">
                <Box className="max-w-4xl" flex="1" overflowY="auto">
                    {children}
                </Box>
            </div>
        </Flex>
    );
}
