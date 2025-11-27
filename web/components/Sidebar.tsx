"use client"

import { Box, VStack, HStack, Text, Icon, Button, Separator } from "@chakra-ui/react";
import { FiLayout, FiTrello, FiHome, FiPlus } from "react-icons/fi";
import Link from "next/link";
import { usePathname } from "next/navigation";

const NavItem = ({ icon, label, href, isActive }: { icon: any, label: string, href: string, isActive?: boolean }) => {
    return (
        <Link href={href} style={{ width: '100%' }}>
            <Button
                variant={isActive ? "subtle" : "ghost"}
                justifyContent="flex-start"
                width="full"
                colorPalette={isActive ? "blue" : "gray"}
                size="sm"
            >
                <Icon as={icon} mr={2} />
                <Text fontWeight={isActive ? "bold" : "normal"}>{label}</Text>
            </Button>
        </Link>
    );
};

export default function Sidebar() {
    const pathname = usePathname();

    return (
        <Box w="260px" h="calc(100vh - 64px)" borderRightWidth="1px" borderColor="border.muted" py={4} px={3} display={{ base: "none", md: "block" }}>
            <VStack align="stretch" gap={1}>
                <NavItem icon={FiTrello} label="Boards" href="/boards" isActive={pathname === "/boards"} />
                <NavItem icon={FiLayout} label="Templates" href="/templates" isActive={pathname === "/templates"} />
                <NavItem icon={FiHome} label="Home" href="/" isActive={pathname === "/"} />

                <Separator my={2} />

                <HStack justify="space-between" px={2} py={1}>
                    <Text fontSize="xs" fontWeight="bold" color="fg.muted">Workspaces</Text>
                    <Icon as={FiPlus} boxSize={3} color="fg.muted" cursor="pointer" />
                </HStack>

                {/* Placeholder for workspaces */}
                <Button variant="ghost" justifyContent="flex-start" width="full" size="sm">
                    <Box boxSize={5} bg="green.500" borderRadius="sm" mr={2} display="flex" alignItems="center" justifyContent="center">
                        <Text fontSize="xs" color="white" fontWeight="bold">N</Text>
                    </Box>
                    <Text>Nexus Workspace</Text>
                </Button>
            </VStack>
        </Box>
    );
}
