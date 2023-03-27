import {useNavigate} from "react-router-dom";

import useStore from "../store";
import {toast} from "react-toastify";
import {useEffect, useState, } from "react";

interface EthPositions {
    positions: EPosition[];
}
interface EPosition {
    ID: string;
    userID: string;
    amount: string;
    currency: string;
    value: string;
    created_at: string;
    updatedAt: string;
}


const EthPosition = () => {
    const [ethPositions, setData] = useState<EPosition[]>([]);
    const [currentPage, setCurrentPage] = useState(1);
    const [itemsPerPage, setItemsPerPage] = useState(10);
    const navigate = useNavigate();
    const store = useStore();
    const fetchPositions = async () => {
        try {
            store.setRequestLoading(true);
            const VITE_SERVER_ENDPOINT = import.meta.env.VITE_SERVER_ENDPOINT;
            const response = await fetch(`${VITE_SERVER_ENDPOINT}/api/eth/positions`, {
                credentials: "include",
            });
            if (!response.ok) {
                throw await response.json();
            }
            const data = await response.json();
            console.log("JSON response:", data); // log the response


            console.log(data)
            store.setRequestLoading(false);

            setData(data);
        }catch (error: any) {
            store.setRequestLoading(false);
            if (error.error) {
                error.error.forEach((err: any) => {
                    toast.error(err.message, {
                        position: "top-right",
                    });
                });
                return;
            }
            const resMessage =
                (error.response &&
                    error.response.data &&
                    error.response.data.message) ||
                error.message ||
                error.toString();

            if (error?.message === "You are not logged in"){
                navigate("/login");
            }

            toast.error(resMessage, {
                position: "top-right",
            });
        }
    };
    useEffect(() => {
        fetchPositions();
    }, []);
    const getPaginatedItems = () => {
        const startIndex = (currentPage - 1) * itemsPerPage;
        const endIndex = startIndex + itemsPerPage;
        return ethPositions.slice(startIndex, endIndex);
    };

    return (
        <section className="bg-ct-blue-50  min-h-screen pt-40">
            <div className="max-w-4xl mx-auto bg-ct-dark-100 rounded-md h-[20rem] flex justify-center items-center">
                <div>
                    <p className="text-5xl text-center font-semibold">ETH Positions Page</p>
                    {!ethPositions ? (
                        <p>Loading...</p>
                    ) : (
                        <table className="border-collapse border-2 border-gray-500">
                            <thead>
                            <tr>
                                <th className="p-2 border-collapse border-2 border-gray-500">ID</th>

                                <th className="p-2 border-collapse border-2 border-gray-500">Currency</th>
                                <th className="p-2 border-collapse border-2 border-gray-500">Amount</th>
                                <th className="p-2 border-collapse border-2 border-gray-500">Created</th>

                            </tr>
                            </thead>
                            <tbody>
                            {getPaginatedItems().map((position) => (
                                <tr key={position.ID}>
                                    <td className="p-2 border-collapse border-2 border-gray-500">{position.ID}</td>

                                    <td className="p-2 border-collapse border-2 border-gray-500">{position.currency}</td>
                                    <td className="p-2 border-collapse border-2 border-gray-500">{position.value}</td>
                                    <td className="p-2 border-collapse border-2 border-gray-500">{position.created_at}</td>

                                </tr>
                            ))}
                            </tbody>

                        </table>
                    )}
                    <div className="flex justify-center items-center mt-4">
                        <button
                            className="px-4 py-2 bg-gray-500 text-white rounded-md mr-4"
                            onClick={() => setCurrentPage(currentPage - 1)}
                            disabled={currentPage === 1}
                        >
                            Previous
                        </button>
                        <button
                            className="px-4 py-2 bg-gray-500 text-white rounded-md"
                            onClick={() => setCurrentPage(currentPage + 1)}
                            disabled={currentPage === Math.ceil(ethPositions.length / itemsPerPage)}
                        >
                            Next
                        </button>
                    </div>

                </div>
            </div>
        </section>
    )
}

export default EthPosition;