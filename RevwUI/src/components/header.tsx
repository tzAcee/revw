import { Card, Dot } from "@geist-ui/core";
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
		<>
            <Card shadow style={{ margin: "10px 10px 10px auto", width: "fit-content", display: "flex"}}>
                <Dot type={isAlive ? "success" : "error"}>
                        {isAlive ? "API alive" : "API unreachable"}
                </Dot>
            </Card>
		</>
	);
}