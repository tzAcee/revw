import { Card, Dot, Text } from "@geist-ui/core";
import { useEffect, useState } from "preact/hooks";
import { revwAPIService } from "../services/revw-api.service";

export function Header() {
    const [isAlive, setIsAlive] = useState(false);

    useEffect(() => {
        async function checkAlive() {
            let revwService = new revwAPIService();
            let aliveStatus = await revwService.Alive();
            setIsAlive(aliveStatus);
        }

        checkAlive();
    }, []);

	return (
        <span style={{ display: "flex", flexDirection: "row" }}>
            <Text h2 style={{margin: "15px 15px 0px 15px"}}>📝 revw</Text>

            <Card shadow style={{ margin: "15px 15px 0px auto", width: "fit-content", display: "flex" }}>
                <Dot type={isAlive ? "success" : "error"}>
                    {isAlive ? "API alive" : "API unreachable"}
                </Dot>
            </Card>
        </span>
	);
}