"use client"

import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { Button, Stack, Heading, Text, Separator, Flex } from "@chakra-ui/react";
import { FcGoogle } from "react-icons/fc";
import Link from "next/link";

import FormCard from "./FormCard";
import FormField from "./FormField";
import { useRouter } from "next/navigation";
import * as userService from "../lib/services/user";
import { toaster } from "@/components/ui/toaster";
import { AxiosError } from "axios";
import { useUserStore } from "@/lib/store/auth";

const schema = z.object({
    email: z.email({ message: "Invalid email address" }),
    password: z.string().min(6, { message: "Password must be at least 6 characters" }),
});

type FormValues = z.infer<typeof schema>;

export default function SigninForm() {
    const {
        register,
        handleSubmit,
        formState: { errors, isSubmitting },
    } = useForm<FormValues>({
        resolver: zodResolver(schema),
    });

    const router = useRouter();
    const userStore = useUserStore();

    const onSubmit = async (data: FormValues) => {
        const promise = userService.login(data);

        toaster.promise(promise, {
            loading: {
                title: "Logging in...",
                description: "Please wait while we log you in",
            },
            success: {
                title: "Logged in!",
                description: "You have been logged in successfully.",
            },
            error: (err) => ({
                title: "Login failed",
                description: err instanceof AxiosError
                    ? err.response?.data?.message || "Something went wrong"
                    : "An unexpected error occurred",
            }),
        });

        try {
            const response = await promise;
            console.log("Login response:", response);
            userStore.setUser(response.data.user);
            router.push("/");
        } catch (error) {
            userStore.clearUser();
            console.log("Login error:", error);
        }
    };

    return (
        <FormCard>
            <Stack gap={6} align="center" w="full">
                <Stack gap={1} align="center">
                    <Heading size="3xl" fontWeight="bold">Nexus</Heading>
                    <Text color="fg.muted" fontSize="lg">Log in to continue</Text>
                </Stack>

                <form onSubmit={handleSubmit(onSubmit)} style={{ width: "100%" }}>
                    <Stack gap={4}>
                        <FormField
                            label="Email"
                            placeholder="Enter your email"
                            type="email"
                            register={register("email")}
                            error={errors.email}
                        />
                        <FormField
                            label="Password"
                            placeholder="Enter your password"
                            type="password"
                            register={register("password")}
                            error={errors.password}
                        />
                        <Button
                            type="submit"
                            colorPalette="blue"
                            loading={isSubmitting}
                            width="full"
                        >
                            Continue
                        </Button>
                    </Stack>
                </form>

                <Flex align="center" w="full" gap={2}>
                    <Separator flex="1" />
                    <Text fontSize="sm" color="fg.muted" whiteSpace="nowrap">
                        Or continue with
                    </Text>
                    <Separator flex="1" />
                </Flex>

                <Button variant="outline" width="full">
                    <FcGoogle />
                    <Text>Google</Text>
                </Button>

                <Text fontSize="sm" color="fg.muted">
                    Don't have an account?{" "}
                    <Link href="/signup" style={{ color: "var(--chakra-colors-blue-500)", fontWeight: "bold" }}>
                        Create one
                    </Link>
                </Text>
            </Stack>
        </FormCard>
    );
}