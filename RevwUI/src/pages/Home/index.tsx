import { Tabs } from "@geist-ui/core";
import { Request } from "./request";
import { Review } from "./review";


export function Home() {
	return (
		<Tabs initialValue="1" align="center" leftSpace={0}>
			<>
			<Tabs.Item label="Request" value="1"><Request></Request></Tabs.Item>
			<Tabs.Item label="Review" value="2"><Review></Review></Tabs.Item>
			</>
		</Tabs>
	);
}

