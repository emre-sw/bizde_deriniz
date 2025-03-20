# Usecases Diagram

## ?? Kullanýcý Ýþlemleri (YUNUS EMRE KESELÝ)

### 1. Kayýt Olma

**Actors:** User, System  
**Steps:**

1. Kullanýcý e-posta ve þifre girerek kayýt olmak ister.
2. Sistem, e-postanýn daha önce kullanýlýp kullanýlmadýðýný kontrol eder.
   - Eðer e-posta kullanýlýyorsa hata döner.
   - Eðer þifre çok zayýfsa hata döner.
3. Sistem, doðrulama e-postasý gönderir.
4. Kullanýcý, doðrulama kodunu girer.
   - Kod hatalýysa hata döner.
   - Kod süresi dolmuþsa hata döner.
5. Kullanýcý hesabý oluþturulur ve giriþ yapýlýr.

### 2. Giriþ Yapma

**Actors:** User, System  
**Steps:**

1. Kullanýcý, e-posta ve þifre ile giriþ yapmak ister.
2. Sistem, bilgilerin doðruluðunu kontrol eder.
   - E-posta yanlýþsa hata döner.
   - Þifre yanlýþsa hata döner.
   - 5 kez hatalý giriþ olursa sistem geçici olarak giriþ engeller.
3. Kullanýcý giriþ yapar.

### 3. Çýkýþ Yapma

**Actors:** User, System  
**Steps:**

1. Kullanýcý "Çýkýþ Yap" butonuna basar.
2. Sistem, kullanýcýyý oturumdan çýkarýr.
3. Kullanýcý ana sayfaya yönlendirilir.

### 4. Þifremi Unuttum

**Actors:** User, System  
**Steps:**

1. Kullanýcý "Þifremi Unuttum" seçeneðine týklar.
2. Kullanýcý, e-posta adresini girer.
3. Sistem, e-postaya doðrulama kodu gönderir.
4. Kullanýcý doðrulama kodunu girer.
5. Kullanýcý yeni þifre belirler.
6. Kullanýcý giriþ yapar.

---

## ?? Kullanýcý Bilgileri (YUNUS ÝNAL)

### 1. Adres Ekleme

**Actors:** User, System  
**Steps:**

1. Kullanýcý, yeni bir adres eklemek ister.
2. Kullanýcý adres bilgilerini girer (þehir, ilçe, açýk adres vb.).
3. Sistem adresi kaydeder.

### 2. Adres Kayýt Etme

**Actors:** User, System  
**Steps:**

1. Kullanýcý mevcut bir adresi kayýtlý adres olarak iþaretler.
2. Sistem adresi varsayýlan olarak belirler.

### 3. Adres Çýkarma

**Actors:** User, System  
**Steps:**

1. Kullanýcý bir adresi silmek ister.
2. Sistem, silme iþlemini onaylamasýný ister.
3. Kullanýcý onay verirse adres kaldýrýlýr.

### 4. Kiþisel Bilgiler (Telefon, Mail) Ekleme

**Actors:** User, System  
**Steps:**

1. Kullanýcý, telefon numarasý veya e-posta adresi eklemek ister.
2. Kullanýcý yeni bilgileri girer.
3. Sistem doðrulama kodu gönderir.
4. Kullanýcý kodu girer ve doðrulama saðlanýr.
5. Bilgiler güncellenir.

---

## ?? Sipariþ Süreci (YUNUS ÝNAL)

### 1. Sepete Ekleme

**Actors:** User, System  
**Steps:**

1. Kullanýcý bir ürünü sepete ekler.
2. Sistem, ürünün stokta olup olmadýðýný kontrol eder.
3. Sistem, ürünü sepete ekler.

### 2. Sepetten Çýkartma

**Actors:** User, System  
**Steps:**

1. Kullanýcý bir ürünü sepetten kaldýrmak ister.
2. Sistem ürünü sepetten çýkarýr.

### 3. Sipariþe Özel Not Ekleme

**Actors:** User, System  
**Steps:**

1. Kullanýcý sipariþ sýrasýnda bir not eklemek ister.
2. Kullanýcý özel notunu girer.
3. Sistem, notu sipariþle birlikte kaydeder.

### 4. Sipariþ Zamaný Planlama

**Actors:** User, System  
**Steps:**

1. Kullanýcý, sipariþin belirli bir zamanda teslim edilmesini ister.
2. Kullanýcý teslimat zamanýný seçer.
3. Sistem sipariþi planlanan zamana göre iþler.

### 5. Tamamlayýcý Ürün Önerileri

**Actors:** User, System  
**Steps:**

1. Kullanýcý bir ürün seçtiðinde, sistem ona ilgili ürünleri önerir.
2. Kullanýcý, önerilen ürünlerden birini veya birkaçýný sepete ekleyebilir.

---

## ?? Arama ve Sýralama Use Cases (YUNUS EMRE KESELÝ)

### 1. Yemek Çeþidine Göre Filtreleme

**Actors:** User, System  
**Steps:**

1. Kullanýcý yemek arama sayfasýna gider.
2. Kullanýcý “Yemek Çeþidi” filtresini açar.
3. Kullanýcý yemek türünü seçer.
4. Sistem uygun restoranlarý listeler.

### 2. Minimum Sepet Tutarýna Göre Filtreleme

**Actors:** User, System  
**Steps:**

1. Kullanýcý minimum sepet tutarý filtresini açar.
2. Kullanýcý belirli bir tutar girer.
3. Sistem belirlenen tutarýn altýndaki restoranlarý listeden çýkarýr.

### 3. Restoran Puanýna Göre Filtreleme

**Actors:** User, System  
**Steps:**

1. Kullanýcý restoran puaný filtresini açar.
2. Kullanýcý belirli bir minimum puan girer.
3. Sistem düþük puanlý restoranlarý listeden çýkarýr.

### 4. Ortalama Varýþ Süresine Göre Filtreleme

**Actors:** User, System  
**Steps:**

1. Kullanýcý varýþ süresi filtresini açar.
2. Kullanýcý belirli bir maksimum süre girer.
3. Sistem belirlenen sürenin üzerindeki restoranlarý listeden çýkarýr.

### 5. Restoranlarý Sýralama

**Actors:** User, System  
**Steps:**

1. Kullanýcý sýralama seçeneðini açar.
2. Kullanýcý sýralama kriterini seçer (teslimat süresi, mesafe, puan).
3. Sistem, restoranlarý seçilen kritere göre sýralar.

---

## ?? Bildirimler Use Cases (YUNUS EMRE KESELÝ)

### 1. Restoraný Favorilere Ekleme

**Actors:** User, System  
**Steps:**

1. Kullanýcý restoran detay sayfasýna girer.
2. Kullanýcý “Favorilere Ekle” butonuna basar.
3. Sistem restoraný favori listesine ekler.

### 2. Yemek Türünü Favorilere Ekleme

**Actors:** User, System  
**Steps:**

1. Kullanýcý yemek türünü seçer.
2. Kullanýcý favorilere ekler.
3. Sistem, seçilen yemek türünü favorilere kaydeder.

### 3. Ýndirim Olduðunda Bildirim Gönderme

**Actors:** User, System  
**Steps:**

1. Kullanýcý restoraný veya yemek türünü favorilere ekler.
2. Sistem, indirim olduðunda bildirim gönderir.

### 4. Mevcut Kampanyalarý Gösterme

**Actors:** User, System  
**Steps:**

1. Kullanýcý kampanyalar sekmesine girer.
2. Sistem, güncel kampanyalarý gösterir.
