import Image from "next/image";
import Link from "next/link";

export default function Home() {
  return (
    <div>
      hello
      <Link href={"/orders"}>Orders</Link>
    </div>
  );
}
