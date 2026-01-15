import { Box, Button, Popover, Tooltip } from "@chakra-ui/react";
import { useRef } from "react";

interface CardActionButtonProps {
    icon: React.ReactNode;
    text: string
    portal?: React.ReactNode;
    isHidden?: boolean;
}

const CardActionButton = ({ icon, text, portal, isHidden = true }: CardActionButtonProps) => {
    return (
        <Popover.Root positioning={{ placement: "bottom-start", flip: false, overlap: true }} lazyMount unmountOnExit>
            <Popover.Trigger asChild>
                <Box>
                    <Button size="xs" variant={"outline"} _hover={{ bg: "gray.700" }} fontSize={"sm"} fontWeight={"medium"} hidden={!isHidden}>
                        {icon}
                        {text}
                    </Button>
                </Box>
            </Popover.Trigger>
            {portal}
        </Popover.Root>

    )
}

export default CardActionButton;