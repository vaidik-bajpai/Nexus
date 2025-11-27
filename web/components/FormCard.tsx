import { Box, Card } from "@chakra-ui/react"

const FormCard = ({ children }: { children: React.ReactNode }) => {
    return (
        <Card.Root maxW="sm" w="full" variant="elevated">
            <Card.Body>
                <Box className="flex flex-col gap-4">
                    {children}
                </Box>
            </Card.Body>
        </Card.Root>
    )
}

export default FormCard