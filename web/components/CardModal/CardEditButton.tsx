import { Button, Popover } from "@chakra-ui/react"
import { Tooltip } from "@/components/ui/tooltip"
import React from "react"

interface CardEditButtonProps {
    icon: React.ReactNode
    tooltipText?: string
    onClick?: () => void
    portal?: React.ReactNode
}

export default function CardEditButtons({ icon, tooltipText, onClick, portal }: CardEditButtonProps) {
    const button = (
        <Button size="xs" variant={"ghost"} w={"fit-content"} px={2} onClick={onClick} borderRadius={"full"} _hover={{ bg: "gray.700" }}>
            {icon}
        </Button>
    )

    if (portal) {
        return (
            <Popover.Root positioning={{ placement: "bottom-start" }}>
                {tooltipText ? (
                    <Tooltip
                        content={tooltipText}
                        positioning={{ placement: "top", offset: { mainAxis: 4, crossAxis: 4 } }}
                        contentProps={{ css: { "--tooltip-bg": "colors.gray.300" } }}
                    >
                        <span style={{ display: "inline-block" }}>
                            <Popover.Trigger asChild>
                                {button}
                            </Popover.Trigger>
                        </span>
                    </Tooltip>
                ) : (
                    <Popover.Trigger asChild>
                        {button}
                    </Popover.Trigger>
                )}
                {portal}
            </Popover.Root>
        )
    }

    if (tooltipText) {
        return (
            <Tooltip
                content={tooltipText}
                positioning={{ placement: "top", offset: { mainAxis: 4, crossAxis: 4 } }}
                contentProps={{ css: { "--tooltip-bg": "colors.gray.300" } }}
            >
                {button}
            </Tooltip>
        )
    }

    return button
}