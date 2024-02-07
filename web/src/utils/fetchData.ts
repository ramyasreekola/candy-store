import { Purchase } from "../types/type";

export const fetchCandyData = async (
  setPurchase: (data: Purchase[]) => void,
  setError: (error: string) => void
) => {
  fetch("/api/candy")
    .then((response) => {
      if (!response.ok) {
        throw new Error(`Error: ${response.statusText}`);
      }
      return response.json();
    })
    .then((data: Purchase[] | void) => {
      if (!data) return;

      const purchasesWithIds = data.map((item, index) => ({
        ...item,
        id: index,
      }));
      setPurchase(purchasesWithIds);
    })
    .catch((error) => setError(error.message));
};
