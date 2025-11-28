import { Flex, Icon, Menu, Text } from "@chakra-ui/react";

interface MenuItemProps {
    logo: React.ReactNode;
    action: string
    title: string
    description: string
}

const MenuItem = ({ logo, action, title, description }: MenuItemProps) => {
    return (
        <Menu.Item value={action}>
            <Flex direction={"column"} gap={1}>
                <Flex align="center" gap={2}>
                    <Icon>
                        {logo}
                    </Icon>
                    <Text fontSize={"sm"}>{title}</Text>
                </Flex>
                <Text fontSize={"xs"} color="gray.500">{description}</Text>
            </Flex>
        </Menu.Item>
    )
}

export default MenuItem;