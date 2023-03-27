import {useNavigate} from "react-router-dom";
import useStore from "../store";
import {Price} from "../store/types";
import {toast} from "react-toastify";
import {useEffect, useState} from "react";

const ETHPage = () => {
    const [currentPrice, setCurrentPrice] = useState( {}as Price)
    const navigate = useNavigate();
    const store = useStore();
    const fetchPrice = async () => {
        try {
            store.setRequestLoading(true);
            const VITE_SERVER_ENDPOINT = import.meta.env.VITE_SERVER_ENDPOINT;
            const response = await fetch(`${VITE_SERVER_ENDPOINT}/api/eth/current-price`, {
                credentials: "include",
            });
            if (!response.ok) {
                throw await response.json();
            }
            const data = await response.json();
            const current_price = data as Price;
            console.log(data)
            store.setRequestLoading(false);
            setCurrentPrice(current_price)
            console.log(current_price);
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

            if (error?.message === "You are not logged in") {
                navigate("/login");
            }

            toast.error(resMessage, {
                position: "top-right",
            });
        }
    };
    useEffect(() => {
        fetchPrice();
    }, []);

    return (
        <section className="bg-ct-blue-600  min-h-screen pt-20">
            <div className="max-w-4xl mx-auto bg-ct-dark-100 rounded-md h-[20rem] flex justify-center items-center">
                <div>
                    <p className="text-5xl text-center font-semibold">Current Price Page</p>
                    {!currentPrice ? (
                        <p>Loading...</p>
                    ) : (
                        <div className="flex items-center gap-8">
                            <div>
                                <img
                                    src = {currentPrice.photo}
                                    className="max-h-36"
                                    alt={`profile photo of ${currentPrice.photo}`}
                                />
                            </div>
                            <div className="mt-8">
                                <p className="mb-3">Currency: {currentPrice.currency}</p>
                                <p className="mb-3">Value: {currentPrice.value}</p>

                            </div>
                        </div>
                    )}
                </div>
            </div>
        </section>
    )


}

export default ETHPage;