"use client";

import { useEffect, useRef } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { useUserStore } from "@/lib/store/auth";
import apiClient from "@/lib/apiClient";
import { Center, Spinner, Text, VStack } from "@chakra-ui/react";

export default function OAuthCallbackPage() {
    const router = useRouter();
    const searchParams = useSearchParams();
    const setUser = useUserStore((state) => state.setUser);
    const processedRef = useRef(false);

    useEffect(() => {
        if (processedRef.current) return;

        const accessToken = searchParams.get("accessToken");
        const refreshToken = searchParams.get("refreshToken");

        if (accessToken && refreshToken) {
            processedRef.current = true;
            localStorage.setItem("accessToken", accessToken);
            localStorage.setItem("refreshToken", refreshToken);

            apiClient.get("/users/me")
                .then((response) => {
                    setUser(response.data.data);
                    router.push("/boards");
                })
                .catch((error) => {
                    console.error("Failed to fetch user details", error);
                    router.push("/login?error=oauth_failed");
                });
        } else {
            const timeout = setTimeout(() => {
                if (!processedRef.current) {
                    router.push("/login?error=missing_tokens");
                }
            }, 1000);
            return () => clearTimeout(timeout);
        }
    }, [searchParams, router, setUser]);

    return (
        <Center h="100vh">
            <VStack gap={4}>
                <Spinner size="xl" color="blue.500" />
                <Text>Authenticating...</Text>
            </VStack>
        </Center>
    );
}
