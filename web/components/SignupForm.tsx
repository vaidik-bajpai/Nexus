"use client"

import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { Button, Stack, Heading, Text, Separator, Flex } from "@chakra-ui/react";
import { FcGoogle } from "react-icons/fc";
import Link from "next/link";
import * as userService from "../lib/services/user";

import FormCard from "./FormCard";
import FormField from "./FormField";
import { AxiosError } from "axios";
import { toaster } from "@/components/ui/toaster";
import { useRouter } from "next/navigation";

const schema = z.object({
    username: z.string().min(3, { message: "Username must be at least 3 characters" }),
    email: z.string().email({ message: "Invalid email address" }),
    password: z.string().min(6, { message: "Password must be at least 6 characters" }),
});

type FormValues = z.infer<typeof schema>;

export default function SignupForm() {
    const {
        register,
        handleSubmit,
        formState: { errors, isSubmitting },
    } = useForm<FormValues>({
        resolver: zodResolver(schema),
    });

    const router = useRouter();
    const onSubmit = async (data: FormValues) => {
        const promise = userService.register(data);

        toaster.promise(promise, {
            loading: {
                title: "Creating account...",
                description: "Please wait while we set up your account",
            },
            success: {
                title: "Account created!",
                description: "We've sent you a verification email.",
            },
            error: (err) => ({
                title: "Registration failed",
                description: err instanceof AxiosError
                    ? err.response?.data?.message || "Something went wrong"
                    : "An unexpected error occurred",
            }),
        });

        try {
            const response = await promise;
            console.log("Registration response:", response);
            router.push("/login");
        } catch (error) {
            console.log("Registration error:", error);
        }
    };

    return (
        <FormCard>
            <Stack gap={6} align="center" w="full">
                <Stack gap={1} align="center">
                    <Heading size="3xl" fontWeight="bold">Nexus</Heading>
                    <Text color="fg.muted" fontSize="lg">Sign up for your account</Text>
                </Stack>

                <form onSubmit={handleSubmit(onSubmit)} style={{ width: "100%" }}>
                    <Stack gap={4}>
                        <FormField
                            label="Username"
                            placeholder="Enter your username"
                            type="text"
                            register={register("username")}
                            error={errors.username}
                        />
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
                            Sign Up
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

                <Button
                    variant="outline"
                    width="full"
                    onClick={() => {
                        window.location.href = `${process.env.NEXT_PUBLIC_API_BASE_URL}/users/google`;
                    }}
                >
                    <FcGoogle />
                    <Text>Google</Text>
                </Button>

                <Text fontSize="sm" color="fg.muted">
                    Already have an account?{" "}
                    <Link href="/login" style={{ color: "var(--chakra-colors-blue-500)", fontWeight: "bold" }}>
                        Log in
                    </Link>
                </Text>
            </Stack>
        </FormCard>
    );
}
