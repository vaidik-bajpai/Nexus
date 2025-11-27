"use client"

import { useUserStore } from "@/lib/store/auth";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

const Layout = ({ children }: { children: React.ReactNode }) => {
    const user = useUserStore((state) => state.user);
    const router = useRouter();

    useEffect(() => {
        if (user) {
            router.push("/");
        }
    }, [user, router])

    if (user) {
        return null
    }

    return <>{children}</>;
};

export default Layout;