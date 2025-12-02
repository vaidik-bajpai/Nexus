import { Button, Popover, Tooltip } from "@chakra-ui/react";

interface CardActionButtonProps {
    icon: React.ReactNode;
    text: string
    portal?: React.ReactNode;
}

const CardActionButton = ({ icon, text, portal }: CardActionButtonProps) => {
    return (
        <Popover.Root positioning={{ placement: "bottom-start" }} lazyMount unmountOnExit>
            <span style={{ display: "inline-block" }}>
                <Popover.Trigger asChild>
                    <Button size="xs" variant={"outline"} _hover={{ bg: "gray.700" }} fontSize={"sm"} fontWeight={"medium"}>
                        {icon}
                        {text}
                    </Button>
                </Popover.Trigger>
            </span>
            {portal}
        </Popover.Root>

    )
}

export default CardActionButton;