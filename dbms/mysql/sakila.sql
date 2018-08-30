use sakila;

SELECT 
    count(distinct rental_id), count(rental_id)
FROM
    sakila.payment;

SELECT 
    rental_id, customer_id
FROM
    sakila.payment
order by 1,2
;

# aggregate of amount for all cust, rental
SELECT 
    rental_id, customer_id, MAX(amount) AS _max
FROM
    sakila.payment
WHERE
    rental_id IS NOT NULL
GROUP BY rental_id , customer_id;


# rental_id for each cutomer having max amount
SELECT 
    p.rental_id, p.customer_id, amount
FROM
    sakila.payment as p
GROUP BY p.rental_id , p.customer_id
HAVING _sum = (SELECT 
        MAX(amount) AS _max
    FROM
        sakila.payment
    WHERE
        rental_id IS NOT NULL
	and p.rental_id = rental_id
    and p.customer_id = customer_id
    GROUP BY rental_id , customer_id);
    
# each cutomer's high rent service and its amount


select * from sakila.film order by 2,5;
select count(distinct title) from sakila.film;

# most contliest film 
SELECT 
    film_id, title, rental_rate
FROM
    sakila.film
group by film_id
# having rental_rate = max(rental_rate)  # wrong
having rental_rate = (select max(rental_rate) from sakila.film)
;

# second most contliest film 
SELECT 
    film_id, title, rental_rate
FROM
    sakila.film
GROUP BY film_id
HAVING rental_rate = (SELECT 
        MAX(rental_rate) AS second_max
    FROM
        sakila.film
    WHERE
        rental_rate < (SELECT 
                MAX(rental_rate) AS _max
            FROM
                sakila.film))
;

# second most contliest film -- using analytical functions TOP/RANK