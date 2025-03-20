# Usecases Diagram

## ?? Kullan�c� ��lemleri (YUNUS EMRE KESEL�)

### 1. Kay�t Olma

**Actors:** User, System  
**Steps:**

1. Kullan�c� e-posta ve �ifre girerek kay�t olmak ister.
2. Sistem, e-postan�n daha �nce kullan�l�p kullan�lmad���n� kontrol eder.
   - E�er e-posta kullan�l�yorsa hata d�ner.
   - E�er �ifre �ok zay�fsa hata d�ner.
3. Sistem, do�rulama e-postas� g�nderir.
4. Kullan�c�, do�rulama kodunu girer.
   - Kod hatal�ysa hata d�ner.
   - Kod s�resi dolmu�sa hata d�ner.
5. Kullan�c� hesab� olu�turulur ve giri� yap�l�r.

### 2. Giri� Yapma

**Actors:** User, System  
**Steps:**

1. Kullan�c�, e-posta ve �ifre ile giri� yapmak ister.
2. Sistem, bilgilerin do�rulu�unu kontrol eder.
   - E-posta yanl��sa hata d�ner.
   - �ifre yanl��sa hata d�ner.
   - 5 kez hatal� giri� olursa sistem ge�ici olarak giri� engeller.
3. Kullan�c� giri� yapar.

### 3. ��k�� Yapma

**Actors:** User, System  
**Steps:**

1. Kullan�c� "��k�� Yap" butonuna basar.
2. Sistem, kullan�c�y� oturumdan ��kar�r.
3. Kullan�c� ana sayfaya y�nlendirilir.

### 4. �ifremi Unuttum

**Actors:** User, System  
**Steps:**

1. Kullan�c� "�ifremi Unuttum" se�ene�ine t�klar.
2. Kullan�c�, e-posta adresini girer.
3. Sistem, e-postaya do�rulama kodu g�nderir.
4. Kullan�c� do�rulama kodunu girer.
5. Kullan�c� yeni �ifre belirler.
6. Kullan�c� giri� yapar.

---

## ?? Kullan�c� Bilgileri (YUNUS �NAL)

### 1. Adres Ekleme

**Actors:** User, System  
**Steps:**

1. Kullan�c�, yeni bir adres eklemek ister.
2. Kullan�c� adres bilgilerini girer (�ehir, il�e, a��k adres vb.).
3. Sistem adresi kaydeder.

### 2. Adres Kay�t Etme

**Actors:** User, System  
**Steps:**

1. Kullan�c� mevcut bir adresi kay�tl� adres olarak i�aretler.
2. Sistem adresi varsay�lan olarak belirler.

### 3. Adres ��karma

**Actors:** User, System  
**Steps:**

1. Kullan�c� bir adresi silmek ister.
2. Sistem, silme i�lemini onaylamas�n� ister.
3. Kullan�c� onay verirse adres kald�r�l�r.

### 4. Ki�isel Bilgiler (Telefon, Mail) Ekleme

**Actors:** User, System  
**Steps:**

1. Kullan�c�, telefon numaras� veya e-posta adresi eklemek ister.
2. Kullan�c� yeni bilgileri girer.
3. Sistem do�rulama kodu g�nderir.
4. Kullan�c� kodu girer ve do�rulama sa�lan�r.
5. Bilgiler g�ncellenir.

---

## ?? Sipari� S�reci (YUNUS �NAL)

### 1. Sepete Ekleme

**Actors:** User, System  
**Steps:**

1. Kullan�c� bir �r�n� sepete ekler.
2. Sistem, �r�n�n stokta olup olmad���n� kontrol eder.
3. Sistem, �r�n� sepete ekler.

### 2. Sepetten ��kartma

**Actors:** User, System  
**Steps:**

1. Kullan�c� bir �r�n� sepetten kald�rmak ister.
2. Sistem �r�n� sepetten ��kar�r.

### 3. Sipari�e �zel Not Ekleme

**Actors:** User, System  
**Steps:**

1. Kullan�c� sipari� s�ras�nda bir not eklemek ister.
2. Kullan�c� �zel notunu girer.
3. Sistem, notu sipari�le birlikte kaydeder.

### 4. Sipari� Zaman� Planlama

**Actors:** User, System  
**Steps:**

1. Kullan�c�, sipari�in belirli bir zamanda teslim edilmesini ister.
2. Kullan�c� teslimat zaman�n� se�er.
3. Sistem sipari�i planlanan zamana g�re i�ler.

### 5. Tamamlay�c� �r�n �nerileri

**Actors:** User, System  
**Steps:**

1. Kullan�c� bir �r�n se�ti�inde, sistem ona ilgili �r�nleri �nerir.
2. Kullan�c�, �nerilen �r�nlerden birini veya birka��n� sepete ekleyebilir.

---

## ?? Arama ve S�ralama Use Cases (YUNUS EMRE KESEL�)

### 1. Yemek �e�idine G�re Filtreleme

**Actors:** User, System  
**Steps:**

1. Kullan�c� yemek arama sayfas�na gider.
2. Kullan�c� �Yemek �e�idi� filtresini a�ar.
3. Kullan�c� yemek t�r�n� se�er.
4. Sistem uygun restoranlar� listeler.

### 2. Minimum Sepet Tutar�na G�re Filtreleme

**Actors:** User, System  
**Steps:**

1. Kullan�c� minimum sepet tutar� filtresini a�ar.
2. Kullan�c� belirli bir tutar girer.
3. Sistem belirlenen tutar�n alt�ndaki restoranlar� listeden ��kar�r.

### 3. Restoran Puan�na G�re Filtreleme

**Actors:** User, System  
**Steps:**

1. Kullan�c� restoran puan� filtresini a�ar.
2. Kullan�c� belirli bir minimum puan girer.
3. Sistem d���k puanl� restoranlar� listeden ��kar�r.

### 4. Ortalama Var�� S�resine G�re Filtreleme

**Actors:** User, System  
**Steps:**

1. Kullan�c� var�� s�resi filtresini a�ar.
2. Kullan�c� belirli bir maksimum s�re girer.
3. Sistem belirlenen s�renin �zerindeki restoranlar� listeden ��kar�r.

### 5. Restoranlar� S�ralama

**Actors:** User, System  
**Steps:**

1. Kullan�c� s�ralama se�ene�ini a�ar.
2. Kullan�c� s�ralama kriterini se�er (teslimat s�resi, mesafe, puan).
3. Sistem, restoranlar� se�ilen kritere g�re s�ralar.

---

## ?? Bildirimler Use Cases (YUNUS EMRE KESEL�)

### 1. Restoran� Favorilere Ekleme

**Actors:** User, System  
**Steps:**

1. Kullan�c� restoran detay sayfas�na girer.
2. Kullan�c� �Favorilere Ekle� butonuna basar.
3. Sistem restoran� favori listesine ekler.

### 2. Yemek T�r�n� Favorilere Ekleme

**Actors:** User, System  
**Steps:**

1. Kullan�c� yemek t�r�n� se�er.
2. Kullan�c� favorilere ekler.
3. Sistem, se�ilen yemek t�r�n� favorilere kaydeder.

### 3. �ndirim Oldu�unda Bildirim G�nderme

**Actors:** User, System  
**Steps:**

1. Kullan�c� restoran� veya yemek t�r�n� favorilere ekler.
2. Sistem, indirim oldu�unda bildirim g�nderir.

### 4. Mevcut Kampanyalar� G�sterme

**Actors:** User, System  
**Steps:**

1. Kullan�c� kampanyalar sekmesine girer.
2. Sistem, g�ncel kampanyalar� g�sterir.
