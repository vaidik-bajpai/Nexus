"use client"

import { useUserStore } from "@/lib/store/auth";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function Home() {
  const user = useUserStore((state) => state.user);
  const router = useRouter();


  useEffect(() => {
    if (user == null) {
      router.push("/login");
    }
  }, [user]);

  return (
    <div>
      <h1>welcome {user?.username}</h1>
    </div>
  );
}
