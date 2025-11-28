import { HStack, Text, Icon, Box } from "@chakra-ui/react";
import { FiInbox, FiCalendar, FiTrello, FiLayers } from "react-icons/fi";

export default function FloatingBottomNav() {
    return (
        <Box
            position="fixed"
            bottom={6}
            left="50%"
            transform="translateX(-50%)"
            bg="gray.900"
            color="white"
            px={4}
            py={2}
            borderRadius="full"
            boxShadow="lg"
            zIndex={1000}
            border="1px solid"
            borderColor="whiteAlpha.200"
        >
            <HStack gap={6}>
                <NavItem icon={FiInbox} label="Inbox" />
                <NavItem icon={FiCalendar} label="Planner" />
                <NavItem icon={FiTrello} label="Board" isActive />
                <NavItem icon={FiLayers} label="Switch boards" />
            </HStack>
        </Box>
    );
}

interface NavItemProps {
    icon: React.ElementType;
    label: string;
    isActive?: boolean;
}

function NavItem({ icon, label, isActive }: NavItemProps) {
    return (
        <HStack
            gap={2}
            cursor="pointer"
            color={isActive ? "blue.400" : "gray.400"}
            _hover={{ color: isActive ? "blue.300" : "white" }}
            transition="colors 0.2s"
        >
            <Icon as={icon} />
            <Text fontSize="sm" fontWeight={isActive ? "bold" : "medium"}>
                {label}
            </Text>
        </HStack>
    );
}
