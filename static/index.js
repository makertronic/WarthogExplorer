/*
*****************************************************************************
*
*                              Warthog Explorer
*
*****************************************************************************
* Web : https://Makertronic-yt.com
* copyright 2023
 */

document.addEventListener("DOMContentLoaded", () => {

    // --------------------- PAGING TX --------------------------------------------
    // Récupérer le container des transactions
    const transactionContainer = document.getElementById("wallet_transactions");

    // Récupérer toutes les transactions
    const transactions = transactionContainer.querySelectorAll(".transaction_hist");

    // Nombre de transactions total
    const totalTransactions = transactions.length;

    // Nombre d'éléments par page
    const perPage = 10; 

    // Page courante (démarre à 0)
    let currentPage = 0;
    prevPageButton.disabled = true;

    // Masquer toutes sauf les 10 premières transactions  
    for(let i = perPage; i < totalTransactions; i++){
      transactions[i].style.display = "none";
    }

    // Boutons précédent/suivant
    const prevPage = document.getElementById("prevPageButton");
    const nextPage = document.getElementById("nextPageButton");

    // Ajouter écouteur d'événement sur les boutons
    nextPage.addEventListener("click", () => {
      currentPage++;
      displayTransactions(currentPage);
    });

    prevPage.addEventListener("click", () => {
      currentPage--;
      displayTransactions(currentPage);  
    });

    // Afficher les transactions pour une page donnée
    function displayTransactions(pageNum){

      if(pageNum === 0) {
        prevPageButton.disabled = true;
      } else {
        prevPageButton.disabled = false; 
      }
      
      if(pageNum === Math.ceil(totalTransactions / perPage) - 1) {
        nextPageButton.disabled = true;
      } else {
        nextPageButton.disabled = false;
      }

      // Calculer les index de début et fin
      const startIndex = pageNum * perPage;
      const endIndex = startIndex + perPage;

      // Masquer toutes les transactions
      for(let i = 0; i < totalTransactions; i++){
        transactions[i].style.display = "none";
      }
      
      // Afficher seulement celles de la page  
      for(let i = startIndex; i < endIndex; i++){
        transactions[i].style.display = "block";
      }

      // Mettre à jour le texte d'information sur la page
      pageInfo.innerText = `Page ${pageNum + 1} sur ${Math.ceil(totalTransactions / perPage)}`; 
    }
});